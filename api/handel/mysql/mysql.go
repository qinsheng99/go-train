package mysql

import (
	"gin/api/tools/common"
	"gin/internal/dao/idao/customer"
	"gin/internal/model"
	"gin/library/pool"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	m customer.MysqlImp
}

func NewMysql(m customer.MysqlImp) *Handler {
	return &Handler{m: m}
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
	defer p.Wait()
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
