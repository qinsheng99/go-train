package demo

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-train/api/tools/common"
	httprequest "github.com/qinsheng99/go-train/library/request"
)

const codeurl = "https://tianqi.aminer.cn/api/v2/multilingual_code_generate"

func (h *Handle) Code(c *gin.Context) {
	var codeReq struct {
		Content string
		Lang    string
		N       int
	}

	if err := c.ShouldBindBodyWith(&codeReq, binding.JSON); err != nil {
		common.Failure(c, err)
		return
	}

	var req = struct {
		Prompt    string `json:"prompt"`
		N         int    `json:"n"`
		Lang      string `json:"lang"`
		Apikey    string `json:"apikey"`
		Apisecret string `json:"apisecret"`
	}{Prompt: codeReq.Content, N: codeReq.N, Lang: codeReq.Lang, Apikey: "xx", Apisecret: "xx"}

	bys, err2 := json.Marshal(req)
	if err2 != nil {
		common.Failure(c, err2)
		return
	}
	var codeResp struct {
		Status int
		Result struct {
			OutPut struct {
				Code []string `json:"code"`
			} `json:"output"`
		} `json:"result"`
	}

	data, err := httprequest.Post(codeurl, bys, nil, nil)
	if err != nil {
		common.Failure(c, err)
		return
	}

	var resp = make(map[string]interface{})

	resp["status"] = 200
	resp["msg"] = "success"
	resp["data"] = nil

	err = json.Unmarshal(data, &codeResp)
	if err != nil {
		common.Failure(c, err)
		return
	}

	if codeResp.Status != 0 {
		resp["status"] = -1
		resp["msg"] = "请求失败"
		common.Success(c, resp)
		return
	}

	if len(codeResp.Result.OutPut.Code) == 0 {
		resp["data"] = h.mag(codeReq.Lang)
		common.Success(c, resp)
		return
	}

	resp["data"] = codeResp.Result.OutPut.Code[0]
	common.Success(c, resp)
	return
}
func (h *Handle) mag(lang string) (m string) {
	if strings.EqualFold(lang, "Python") {
		m = `\n# Code generation finished, modify this comment to continue the generation.`
	} else {
		m = `\n// Code generation finished, modify this comment to continue the generation.`
	}
	return
}
