package postgresql

import (
	"strconv"

	"github.com/gin-gonic/gin"
	postgresqlRequest "github.com/qinsheng99/go-train/api/entity/postgresql"
	"github.com/qinsheng99/go-train/api/tools/common"
	"github.com/qinsheng99/go-train/internal/dao/idao/boy"
	"github.com/qinsheng99/go-train/internal/model"
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
	res, err := p.boy.GetBoyAddress("address", s)
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

func (p *Handler) FindArrOne(c *gin.Context) {
	i, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if err != nil {
		common.Failure(c, err)
		return
	}
	arr, err := p.boy.FindArrOne(i)
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, arr)
}

func (p *Handler) FindJson(c *gin.Context) {
	query := c.Query("query")

	data, err := p.boy.FindJson(query, c.Query("flag") == "true")
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, data)
}
