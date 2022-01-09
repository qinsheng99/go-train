package routes

import (
	"gin/api"
	"gin/api/handel/mysql"
	"gin/api/middleware"

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
