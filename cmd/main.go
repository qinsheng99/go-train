package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/routes"
	"github.com/qinsheng99/goWeb/library/config"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/elasticsearch"
	mongoClient "github.com/qinsheng99/goWeb/library/mongo"
	"github.com/qinsheng99/goWeb/library/redisClient"

	//_ "github.com/qinsheng99/goWeb/library/etcd"
	"github.com/qinsheng99/goWeb/library/logger"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	if err := config.Init(false); err != nil {
		panic(err)
	}
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	r := gin.Default()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	bundleDB, err := db.GetBundleDb(config.Conf.MysqlConfig)

	if err != nil {
		fmt.Printf("Mysql connect failed , error is %v\n", err)
		panic(err)
	}
	es, err := elasticsearch.GetES(config.Conf.EsConfig)
	if err != nil {
		fmt.Printf("ES connect failed , error is %v\n", err)
		panic(err)
	}
	redis, err := redisClient.GetRedis(config.Conf.RedisConfig)
	if err != nil {
		fmt.Printf("Redis connect failed , error is %v\n", err)
		panic(err)
	}

	mo, err := mongoClient.InitMongo()
	must(err)

	e, err := api.Init(bundleDB, es, redis, mo)
	must(err)
	routes.Route(e, r)

	must(r.Run(fmt.Sprintf(":%v", config.Conf.Port)))
}
