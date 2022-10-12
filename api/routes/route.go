package routes

import (
	"github.com/qinsheng99/goWeb/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Route(e *api.Entry, c *gin.Engine) {
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	CeShi(e, c)
	Credis(e, c)
	Mysql(e, c)
	Sort(e, c)
	Es(e, c)
	Mon(e, c)
	Demo(e, c)
	Postgresql(e, c)
}
