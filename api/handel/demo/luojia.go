package demo

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api/tools/common"
	httprequest "github.com/qinsheng99/goWeb/library/request"
)

const luojia = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/e3f5d06e-4c1f-4b62-b2f4-123d6c22d244/infer/image"

func (h *Handle) Luojia(c *gin.Context) {
	//  - 桶名：luojianet
	//  - 路径：infer/{user_name}/input.png
	username := c.Query("username")

	body := fmt.Sprintf(`{"user_name":"%s"}`, username)

	t := h.gettoken()
	head := map[string]string{
		"X-Auth-Token": t,
	}

	data, err := httprequest.Post(luojia, []byte(body), head, nil)
	if err != nil {
		common.Failure(c, err)
		return
	}

	var luojiares struct {
		Status      int
		Msg, Result string
	}

	err = json.Unmarshal(data, &luojiares)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, luojiares)
}
