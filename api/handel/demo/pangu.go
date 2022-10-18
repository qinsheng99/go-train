package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api/tools/common"
	httprequest "github.com/qinsheng99/goWeb/library/request"
)

const pu = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/302dc54c-144f-425d-b0e7-767d07f8d180/infer/text"

func (h *Handle) Pangu(c *gin.Context) {
	q := c.Query("question")

	t, err := h.redis.Get(context.Background(), "modelarts-token")
	if len(t) == 0 || err != nil {
		t, err = token()
		if err != nil {
			common.Failure(c, err)
			return
		}
		_, _ = h.redis.Set(context.Background(), "modelarts-token", t, time.Hour*24)
	}
	head := map[string]string{
		"X-Auth-Token": t,
	}
	var data []byte

	data, err = httprequest.Post(pu, []byte(fmt.Sprintf(`{"question":"%s"}`, q)), head)
	if err != nil {
		common.Failure(c, err)
		return
	}

	var pres = struct {
		Result string
	}{}

	err = json.Unmarshal(data, &pres)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, pres)
}
