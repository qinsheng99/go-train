package routes

import (
	"github.com/qinsheng99/goWeb/api"

	"github.com/gin-gonic/gin"
)

func Route(e *api.Entry, c *gin.Engine) {
	CeShi(e, c)
	Credis(e, c)
	Mysql(e, c)
	Sort(e, c)
	Es(e, c)
}
