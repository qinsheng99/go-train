package routes

import (
	"github.com/qinsheng99/go-train/api"
	"github.com/qinsheng99/go-train/api/handel/mysql"
	"github.com/qinsheng99/go-train/api/middleware"

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
}
