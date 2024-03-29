package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-train/api"
	"github.com/qinsheng99/go-train/api/handel/postgresql"
	"github.com/qinsheng99/go-train/api/middleware"
)

func Postgresql(e *api.Entry, c *gin.Engine) {
	groupPostgres := c.Group("/postgres").
		Use(middleware.AuthMiddleware())
	func(m *postgresql.Handler) {
		{
			groupPostgres.POST("/create-data", m.CreateOne)
			groupPostgres.GET("/get-data", m.GetPostgresData)
			groupPostgres.GET("/get-filter-data", m.GetPostgresFilter)
			groupPostgres.GET("/get-one", m.GetPostgresOne)
			groupPostgres.GET("/find-arr", m.FindArrOne)
			groupPostgres.GET("/find-json", m.FindJson)
		}
	}(e.NewPostgreSql)
}
