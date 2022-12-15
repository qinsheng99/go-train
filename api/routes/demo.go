package routes

import (
	"net/http"

	"github.com/qinsheng99/go-train/api"
	"github.com/qinsheng99/go-train/api/handel/demo"
	"github.com/qinsheng99/go-train/api/middleware"

	"github.com/gin-gonic/gin"
)

func Demo(e *api.Entry, c *gin.Engine) {
	c.StaticFS("/images", http.Dir("../images"))
	c.MaxMultipartMemory = 174080
	group := c.Group("/demo").
		Use(middleware.AuthMiddleware())
	func(d *demo.Handle) {
		{
			group.POST("/file", d.File)
			group.POST("/text-image", d.TestImage)
			group.POST("/vqa/:v", d.Vqa)

			group.POST("/pangu", d.Pangu)

			group.GET("/luojia", d.Luojia)

			group.POST("/code", d.Code)
			group.POST("/code2", d.Code2)

			group.POST("/wukong", d.WuKong)

		}
	}(e.NewDemo)
}
