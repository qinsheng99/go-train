package mysql

import (
	"strconv"
	"time"

	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/internal/dao/idao/boy"
	"github.com/qinsheng99/goWeb/internal/dao/idao/customer"
	"github.com/qinsheng99/goWeb/internal/model"
	"github.com/qinsheng99/goWeb/library/pool"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	m   customer.CeshiMysqlImp
	boy boy.Boyimpl
}

func NewMysql(m customer.CeshiMysqlImp, boy boy.Boyimpl) *Handler {
	return &Handler{m: m, boy: boy}
}

func (m *Handler) GetData(c *gin.Context) {
	res, err := m.m.GetCeshiData()
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (m *Handler) IntertData(c *gin.Context) {
	var update = []string{"modify_time"}
	p := pool.NewGoPool(pool.WithMaxLimit(3))
	p.Submit(func() {
		for i := 1; i < 6; i++ {
			var data = &model.Ceshi{
				Id:         i,
				Uri:        "ceshi" + strconv.Itoa(i),
				BackName:   "GP" + strconv.Itoa(i),
				IsDelete:   0,
				CreateTime: time.Now().Unix(),
				ModifyTime: time.Now(),
			}
			err := m.m.InsertData(data, update)
			if err != nil {
				common.Failure(c, err)
				return
			}
		}
	})
	defer p.Close()
	common.Success(c, gin.H{})
}

func (m *Handler) JoinData(c *gin.Context) {
	res, err := m.m.JoinData()
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}

func (m *Handler) GetPostgresData(c *gin.Context) {
	res, err := m.boy.Getlist()
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, res)
}
