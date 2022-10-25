package demo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-train/api/tools/common"
	httprequest "github.com/qinsheng99/go-train/library/request"
)

const luojia = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/e3f5d06e-4c1f-4b62-b2f4-123d6c22d244/infer/image"

var luoc chan string

func init() {
	luoc = make(chan string, 1)
	luoc <- luojia
}
func (h *Handle) Luojia(c *gin.Context) {
	//  - 桶名：luojianet
	//  - 路径：infer/{user_name}/input.png
	select {
	case luourl := <-luoc:
		result, err := h.luojia(c, luourl)
		luoc <- luourl
		if err != nil {
			common.Failure(c, err)
			return
		}
		common.Success(c, result)
	default:
		common.Failure(c, errors.New("busy"))
	}
}

func (h *Handle) luojia(c *gin.Context, url string) (string, error) {
	username := c.Query("username")

	body := fmt.Sprintf(`{"user_name":"%s"}`, username)
	t := h.gettoken()
	head := map[string]string{
		"X-Auth-Token": t,
	}

	data, err := httprequest.Post(url, []byte(body), head, nil)
	if err != nil {
		return "", err
	}

	var luojiares struct {
		Status      int
		Msg, Result string
	}

	err = json.Unmarshal(data, &luojiares)
	if err != nil {
		return "", err
	}
	if luojiares.Status != 200 {
		return "", errors.New("推理失败")
	}

	return luojiares.Result, nil
}
