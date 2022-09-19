package routes

import (
	"net/http"

	"github.com/qinsheng99/goWeb/api"
	"github.com/qinsheng99/goWeb/api/handel/demo"
	"github.com/qinsheng99/goWeb/api/middleware"

	"github.com/gin-gonic/gin"
)

func Demo(e *api.Entry, c *gin.Engine) {
	c.StaticFS("/images", http.Dir("../images"))
	c.MaxMultipartMemory = 8 << 15
	group := c.Group("/demo").
		Use(middleware.AuthMiddleware())
	func(d *demo.Handle) {
		{
			group.POST("/file", d.File)
			group.POST("/text-image", d.TestImage)

		}
	}(e.NewDemo)
}
