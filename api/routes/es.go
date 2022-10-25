package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-train/api"
	esHandle "github.com/qinsheng99/go-train/api/handel/es"
	"github.com/qinsheng99/go-train/api/middleware"
)

func Es(e *api.Entry, c *gin.Engine) {
	group := c.Group("/es").
		Use(middleware.AuthMiddleware())
	func(e *esHandle.EsHandle) {
		{
			group.GET("/refresh", e.Refresh)
			group.GET("/getAllEs", e.GetAllEs)
		}
	}(e.NewEs)
}
