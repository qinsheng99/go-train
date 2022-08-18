package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/handel/mongo"
)

func Mon(e *api.Entry, c *gin.Engine) {
	group := c.Group("/mongo")
	func(e *mongo.Handle) {
		{
			group.POST("/insert-one", e.InsertOne)
		}
	}(e.NewMgo)
}
