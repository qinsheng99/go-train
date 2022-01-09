package routes

import (
	"gin/api"

	"github.com/gin-gonic/gin"
)

func Route(e *api.Entry, c *gin.Engine) {
	CeShi(e, c)
	Credis(e, c)
	Mysql(e, c)
}
