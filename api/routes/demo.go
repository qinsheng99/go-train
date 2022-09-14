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
	group := c.Group("/public").
		Use(middleware.AuthMiddleware())
	func(d *demo.Handle) {
		{
			group.Any("/file", d.File)

		}
	}(e.NewDemo)
}
