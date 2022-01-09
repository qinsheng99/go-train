package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CeFang = "cft-userInfo"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("start", start)
		c.Next()
	}
}
