package routes

import (
	"gin/api"
	"gin/api/handel/redis"
	"gin/api/middleware"

	"github.com/gin-gonic/gin"
)

func Credis(e *api.Entry, c *gin.Engine) {
	group := c.Group("/redis").
		Use(middleware.AuthMiddleware())
	func(h *redis.Handle) {
		{
			group.GET("/set", h.SetR)
			group.GET("/get", h.GetR)
			group.GET("/del", h.DelR)
			group.GET("/exists", h.ExistsR)
			group.GET("/dump", h.Dump)

			group.GET("/mset", h.MSet)
			group.GET("/mget", h.MGet)

			group.POST("/hset", h.Hset)
			group.POST("/hget", h.HGetOrAll)

			group.GET("/lpush", h.Lpush)
			group.GET("/lRange", h.LRange)
			group.GET("/lpop", h.Lpop)
			group.GET("/llen", h.Llen)

			group.GET("/sadd", h.Sadd)
			group.GET("/sMembers", h.SMembers)
			group.GET("/sRandMember", h.SRandMember)

			group.GET("/zadd", h.Zadd)
			group.GET("/zrange", h.Zrange)
			group.GET("/zrank", h.Zrank)

			group.GET("/grpc", h.Grpc)
		}
	}(e.NewH)
}
