package demo

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api/tools/common"
	httprequest "github.com/qinsheng99/goWeb/library/request"
)

const pu = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/302dc54c-144f-425d-b0e7-767d07f8d180/infer/text"

func (h *Handle) Pangu(c *gin.Context) {
	q := c.Query("question")

	t := h.gettoken()
	head := map[string]string{
		"X-Auth-Token": t,
	}
	var (
		data []byte
		err  error
	)

	data, err = httprequest.Post(pu, []byte(fmt.Sprintf(`{"question":"%s"}`, q)), head, nil)
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
	var result string = pres.Result

	if i := strings.IndexByte(pres.Result, '\n'); i > 0 && q == pres.Result[:i] && i+1 < len(pres.Result) {
		result = pres.Result[i+1:]
	}

	common.Success(c, result)
}
