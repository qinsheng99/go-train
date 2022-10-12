package postgresql

import (
	"strconv"

	"github.com/gin-gonic/gin"
	postgresqlRequest "github.com/qinsheng99/goWeb/api/entity/postgresql"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/internal/dao/idao/boy"
	"github.com/qinsheng99/goWeb/internal/model"
)

type Handler struct {
	boy boy.BoyimplService
}

func NewPostgreSql(boy boy.BoyimplService) *Handler {
	return &Handler{boy: boy}
}

func (p *Handler) GetPostgresData(c *gin.Context) {
	res, err := p.boy.GetBoylist()
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (p *Handler) GetPostgresFilter(c *gin.Context) {
	s := c.Query("address")
	res, err := p.boy.GetBoyAddress(s)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (p *Handler) GetPostgresOne(c *gin.Context) {
	i, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	res, err := p.boy.GetBoyOne(i)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (p *Handler) CreateOne(c *gin.Context) {
	var data, err, by = &model.Boy{}, error(nil), postgresqlRequest.Boy{}
	if err = c.ShouldBindJSON(&by); err != nil {
		common.Failure(c, err)
		return
	}

	data, err = p.boy.CreateOne(by)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, data)
}
