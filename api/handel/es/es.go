package esHandle

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-train/api/tools/common"
	"github.com/qinsheng99/go-train/internal/dao/idao"
	"github.com/qinsheng99/go-train/internal/dao/idao/customer"
	"github.com/qinsheng99/go-train/library/pool"
)

type EsHandle struct {
	EsDao idao.EsImp
	m     customer.CeshiMysqlImp
}

func NewEsHandle(EsDao idao.EsImp, m customer.CeshiMysqlImp) *EsHandle {
	return &EsHandle{EsDao: EsDao, m: m}
}

func (es *EsHandle) Refresh(c *gin.Context) {
	begin := time.Now()
	p := pool.NewGoPool(pool.WithMaxLimit(1))
	defer p.Close()

	p.Submit(func() {
		ceshiEs, err := es.m.GetCeshiEsData()
		if err != nil {
			common.Failure(c, err)
			return
		}
		es.EsDao.RefreshCeshi(ceshiEs)
	})
	usedTime := time.Since(begin)
	common.Success(c, fmt.Sprintf("es 数据刷新完成，耗时：%s", usedTime.String()))
}

func (es *EsHandle) GetAllEs(c *gin.Context) {
	data, err := es.EsDao.GetAllEsData()
	if err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, data)
}
