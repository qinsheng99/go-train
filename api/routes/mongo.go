package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/qinsheng99/go-train/api"
	"github.com/qinsheng99/go-train/api/handel/mongo"
	"github.com/qinsheng99/go-train/api/middleware"
)

func Mon(e *api.Entry, c *gin.Engine) {
	group := c.Group("/mongo").
		Use(middleware.AuthMiddleware())
	func(e *mongo.Handle) {
		{
			group.POST("/insert-one", e.InsertOne)
			group.POST("/insert-many", e.InsertMany)
			group.GET("/find", e.Find)
			group.GET("/find-one", e.FindOne)
			group.GET("/update", e.Update)
			group.GET("/push", e.Push)

			group.GET("/insert-wukong", e.InsertWukong)
			group.GET("/find-wukong", e.FindWukong)
			group.POST("/aggregate-wukong", e.AggregateWukong)
		}
	}(e.NewMgo)
}
