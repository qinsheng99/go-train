package main

import (
	"fmt"

	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/routes"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/elasticsearch"
	"github.com/qinsheng99/goWeb/library/redisClient"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	bundleDB, err := db.GetBundleDb()

	if err != nil {
		fmt.Printf("Mysql connect failed , error is %v", err)
		panic(err)
	}
	es, err := elasticsearch.GetES()
	if err != nil {
		fmt.Printf("ES connect failed ,  error is %v", err)
		panic(err)
	}
	conn, err := redisClient.GetRedis()
	if err != nil {
		fmt.Printf("Redis connect failed ,  error is %v", err)
		panic(err)
	}

	e, err := api.Init(bundleDB, es, conn)

	if err != nil {
		panic(err)
	}
	routes.Route(e, r)

	_ = r.Run(":111")
}
