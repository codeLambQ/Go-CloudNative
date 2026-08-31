package data

import (
	"moddy-blog-article/internal/conf"
	"moddy-blog-article/internal/data/ent"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	MysqlClient *ent.Client
	RedisClient *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(logger)
	driver, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		logHelper.Errorf("数据库连接失败: %s", err.Error())
		return nil, nil, err
	}
	mysqlClient := ent.NewClient(ent.Driver(driver))

	redisClient := redis.NewClient(&redis.Options{Addr: c.Redis.Addr})

	data := &Data{
		MysqlClient: mysqlClient,
		RedisClient: redisClient,
	}
	cleanup := func() {
		err := mysqlClient.Close()
		if err != nil {
			logHelper.Errorf("数据库关闭失败: %s", err.Error())
			return
		}
		err = redisClient.Close()
		if err != nil {
			logHelper.Errorf("redis关闭失败: %s", err.Error())
			return
		}
	}
	return data, cleanup, nil
}
