package main

import (
	"fmt"

	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/routes"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/elasticsearch"
	etcd2 "github.com/qinsheng99/goWeb/library/etcd"
	"github.com/qinsheng99/goWeb/library/redisClient"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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

	etcd, err := etcd2.GetEtcd()
	if err != nil {
		fmt.Printf("Etcd connect failed , error is %v\n", err)
		panic(err)
	}

	e, err := api.Init(bundleDB, es, redis, etcd)

	if err != nil {
		panic(err)
	}
	routes.Route(e, r)

	_ = r.Run(":111")
}
