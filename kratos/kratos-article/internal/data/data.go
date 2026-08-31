package data

import (
	"context"
	"fmt"
	"kratos-article/internal/conf"
	"kratos-article/internal/data/ent"
	"kratos-article/internal/data/ent/article"

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
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	loggerHelper := log.NewHelper(logger)
	// 创建 mysql 客户端
	driver, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		loggerHelper.Error("数据库连接失败: " + err.Error())
		return nil, nil, err
	}
	mysqlClient := ent.NewClient(ent.Driver(driver))
	loggerHelper.Info("数据库客户端创建成功")

	err = mysqlClient.Schema.Create(context.Background())

	// 对表进行操作
	// mysqlClient.Article.Create().SetID(1).SetTitle("first title").SetContent("no content").Save(context.Background())

	// 查询表
	articles, _ := mysqlClient.Article.Query().Where(article.TitleEQ("first title")).Only(context.Background())
	fmt.Printf("select: %+v", articles)
	if err != nil {
		loggerHelper.Error("数据库表创建失败: " + err.Error())
		return nil, nil, err
	}

	// 创建 redis 客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
	})
	d := &Data{
		db:  mysqlClient,
		rdb: redisClient,
	}
	cleanup := func() {
		_ = d.db.Close()
		_ = d.rdb.Close()
	}
	return d, cleanup, nil
}
