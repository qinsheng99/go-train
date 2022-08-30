package routes

import (
	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/handel/mysql"
	"github.com/qinsheng99/goWeb/api/middleware"

	"github.com/gin-gonic/gin"
)

func Mysql(e *api.Entry, c *gin.Engine) {
	group := c.Group("/mysql").
		Use(middleware.AuthMiddleware())
	func(m *mysql.Handler) {
		{
			group.GET("/get-data", m.GetData)
			group.GET("/join-data", m.JoinData)
			group.GET("/intert-data", m.IntertData)
		}
	}(e.NewMysql)

	groupPostgres := c.Group("/postgres").
		Use(middleware.AuthMiddleware())
	func(m *mysql.Handler) {
		{
			groupPostgres.GET("/get-data", m.GetPostgresData)
		}
	}(e.NewMysql)
}
