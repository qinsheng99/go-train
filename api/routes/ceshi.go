package routes

import (
	"gin/api"
	"gin/api/handel/ceshi"
	"gin/api/middleware"

	"github.com/gin-gonic/gin"
)

func CeShi(e *api.Entry, c *gin.Engine) {
	group := c.Group("/public").
		Use(middleware.AuthMiddleware())
	func(h *ceshi.Handler) {
		{
			// group.GET("/index", h.CeShi)
			group.POST("/index", h.Index)
			group.GET("/get-es", h.GetEs)
			group.GET("/get-es-list", h.GetEsList)
			group.GET("/delete-es", h.DeleteEs)
			group.GET("/create-es", h.CreateEs)
			group.GET("/get-list", h.GetList)
			group.GET("/drainage-list", h.GetDrainageList)
			group.GET("/http", h.Http)

			group.GET("/li-kou", h.LiKou)

		}
	}(e.NewHandler)
}
