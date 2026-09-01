package data

import (
	"context"
	"log/slog"

	"moddy-blog-article/internal/conf"

	"moddy-blog-article/internal/data/ent"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v3/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	mysqlDB *ent.Client
	redisDB *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger *slog.Logger) (*Data, func(), error) {
	log.SetDefault(logger)
	// 创建 mysql 连接
	driver, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logger.Error("数据库连接失败")
		panic(err)
	}
	// 创建 mysql 客户端
	mysqlClient := ent.NewClient(ent.Driver(driver))
	err = mysqlClient.Schema.Create(context.Background())
	if err != nil {
		logger.Error("mysql 表创建失败")
	}
	logger.Info("mysql 表创建成功")
	redisClient := redis.NewClient(&redis.Options{Addr: c.Redis.Addr})
	cleanup := func() {
		err = mysqlClient.Close()
		if err != nil {
			logger.Error("mysql 关闭失败")
		}
		err = redisClient.Close()
		if err != nil {
			logger.Error("redis 关闭失败")
		}
	}
	logger.Info("数据库初始化成功")

	return &Data{
		mysqlDB: mysqlClient,
		redisDB: redisClient,
	}, cleanup, nil
}
