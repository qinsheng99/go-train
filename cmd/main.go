package main

import (
	"fmt"

	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/routes"
	"github.com/qinsheng99/goWeb/library/config"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/elasticsearch"
	_ "github.com/qinsheng99/goWeb/library/etcd"
	"github.com/qinsheng99/goWeb/library/logger"
	"github.com/qinsheng99/goWeb/library/redisClient"

	"github.com/gin-gonic/gin"
)

func must(err error)  {
	if err != nil {
		panic(err)
	}
}
func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	r := gin.Default()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	bundleDB, err := db.GetBundleDb()

	if err != nil {
		fmt.Printf("Mysql connect failed , error is %v\n", err)
		panic(err)
	}
	es, err := elasticsearch.GetES()
	if err != nil {
		fmt.Printf("ES connect failed , error is %v\n", err)
		panic(err)
	}
	redis, err := redisClient.GetRedis()
	if err != nil {
		fmt.Printf("Redis connect failed , error is %v\n", err)
		panic(err)
	}

	e, err := api.Init(bundleDB, es, redis)
	must(err)
	routes.Route(e, r)

	must(r.Run(fmt.Sprintf(":%v",config.Conf.Port)))
}
