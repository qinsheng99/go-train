package demo

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-train/api/tools/common"
	httprequest "github.com/qinsheng99/go-train/library/request"
)

const wukong = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/b8510b21-004a-4a93-bbae-68b11732f057/infer"

type wukongRequest struct {
	Input string `json:"input_text" binding:"required"`
	User  string `json:"user_name"  binding:"required"`
}

type wukongResponse struct {
	Status int      `json:"status"`
	Output []string `json:"output_image_url"`
	Msg    string   `json:"msg"`
}

func (h *Handle) WuKong(c *gin.Context) {
	var req wukongRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.QueryFailure(c, err)
		return
	}

	token := h.gettoken()
	if len(token) == 0 {
		common.Failure(c, fmt.Errorf("token is empty"))
		return
	}

	bys, err := h.wukong(wukong, token, req)
	if err != nil {
		common.Failure(c, err)
		return
	}

	var wk wukongResponse
	err = json.Unmarshal(bys, &wk)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, wk)
}

func (h *Handle) wukong(url, token string, req wukongRequest) ([]byte, error) {
	var data []byte
	var err error

	head := map[string]string{
		"X-Auth-Token": token,
	}

	data, err = httprequest.Post(url, fmt.Sprintf(`{"user_name":"%s","input_text":"%s"}`, req.User, req.Input), head, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}
