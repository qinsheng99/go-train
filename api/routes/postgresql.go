package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/handel/postgresql"
	"github.com/qinsheng99/goWeb/api/middleware"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/internal/model"
	"github.com/qinsheng99/goWeb/library/db"
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
			groupPostgres.GET("/join-data", func(c *gin.Context) {
				var b struct {
					model.Boy
					Nnme string
				}

				db.GetPostgresqlDb().Raw("select a.*,b.name as nnme from boy a left join newboy b on a.id = b.id where a.id = 1").First(&b)

				common.Success(c, b)
			})
		}
	}(e.NewPostgreSql)
}
