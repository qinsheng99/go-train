package demo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-train/api/tools/common"
	httprequest "github.com/qinsheng99/go-train/library/request"
)

const pu = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/302dc54c-144f-425d-b0e7-767d07f8d180/infer/text"

var puc chan string

func init() {
	puc = make(chan string, 1)
	puc <- pu
}

func (h *Handle) Pangu(c *gin.Context) {
	select {
	case panurl := <-puc:
		result, err := h.pangu(c, panurl)
		puc <- panurl
		if err != nil {
			common.Failure(c, err)
			return
		}
		common.Success(c, result)
	default:
		common.Failure(c, errors.New("busy"))
	}
}

func (h *Handle) pangu(c *gin.Context, url string) (string, error) {
	q := c.Query("question")

	t := h.gettoken()
	head := map[string]string{
		"X-Auth-Token": t,
	}
	var (
		data []byte
		err  error
	)

	data, err = httprequest.Post(url, []byte(fmt.Sprintf(`{"question":"%s"}`, q)), head, nil)
	if err != nil {
		return "", err
	}

	var pres = struct {
		Result string
	}{}

	err = json.Unmarshal(data, &pres)
	if err != nil {
		common.Failure(c, err)
		return "", err
	}
	var result string = pres.Result

	if i := strings.IndexByte(pres.Result, '\n'); i > 0 && q == pres.Result[:i] && i+1 < len(pres.Result) {
		result = pres.Result[i+1:]
	}
	return result, nil
}
