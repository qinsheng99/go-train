package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api"
	sortHandler "github.com/qinsheng99/goWeb/api/handel/sort"
	"github.com/qinsheng99/goWeb/api/middleware"
)

func Sort(e *api.Entry, c *gin.Engine) {
	group := c.Group("/sort").
		Use(middleware.AuthMiddleware())
	func(s *sortHandler.SortHandler) {
		{
			group.GET("/select-sort", s.SelectSort)
			group.GET("/bubbling-sort", s.BubblingSort)
			group.GET("/insert-sort", s.InsertSort)
			group.GET("/shell-sort", s.ShellSort)
			group.GET("/merge-sort", s.MergeSort)
			group.GET("/quick-sort", s.QuickSort)
			group.GET("/count-sort", s.CountSort)
		}
	}(e.NewSort)
}
