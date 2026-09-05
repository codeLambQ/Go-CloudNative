package data

import (
	"context"
	"fmt"
	"log/slog"

	"moddy-blog-article/internal/conf"

	"moddy-blog-article/internal/data/ent"

	"github.com/go-kratos/kratos/v3/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

// Data .
type Data struct {
	PDB     *ent.Client
	redisDB *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger *slog.Logger) (*Data, func(), error) {
	log.SetDefault(logger)
	/* 创建 mysql 连接
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
	*/

	// 创建 postgres 连接
	fmt.Printf("drive: %s  source:%s", c.Database.Driver, c.Database.Source)
	pClient, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logger.Error("数据库连接失败")
		return nil, nil, err
	}

	err = pClient.Schema.Create(context.Background())
	if err != nil {
		logger.Error("pg 表创建失败")
		return nil, nil, err
	}
	logger.Info("pg 表创建成功")

	// 创建数据
	article, err := pClient.Article.Create().SetTitle("第一个").SetContent("文章内容是我不知道是啥").Save(context.Background())
	if err != nil {
		logger.Error("pg 表数据添加失败")
		return nil, nil, err
	}
	logger.Info("数据库创建完成，article = ", article)

	redisClient := redis.NewClient(&redis.Options{Addr: c.Redis.Addr})
	cleanup := func() {
		err = pClient.Close()
		if err != nil {
			logger.Error("pg 关闭失败")
		}
		err = redisClient.Close()
		if err != nil {
			logger.Error("redis 关闭失败")
		}
	}
	logger.Info("数据库初始化成功")

	return &Data{
		PDB:     pClient,
		redisDB: redisClient,
	}, cleanup, nil
}
