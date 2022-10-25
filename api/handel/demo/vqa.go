package demo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	path2 "path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-train/api/entity/demo"
	"github.com/qinsheng99/go-train/api/tools/common"
)

func (h *Handle) vqa2(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		common.Failure(c, err)
		return
	}
	if f.Size >= 200000 {
		common.Failure(c, errors.New("the file size cannot be larger than 200kb"))
		return
	}
	question := c.PostForm("question")
	s := path2.Ext(f.Filename)

	if !strings.Contains(strings.ToLower(s), "png") && !strings.Contains(strings.ToLower(s), "jpg") {
		common.Failure(c, errors.New("image not jpg/png"))
		return
	}
	path := "../images/" + f.Filename
	err = c.SaveUploadedFile(f, path)
	if err != nil {
		common.Failure(c, err)
		return
	}
	var (
		req       *http.Request
		response  *http.Response
		bys       []byte
		imageFile *os.File
		t         string
	)
	t = h.gettoken()

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	file, _ := writer.CreateFormFile("file", f.Filename)
	imageFile, err = os.Open(path)
	if err != nil {
		common.Failure(c, err)
		return
	}
	_, err = io.Copy(file, imageFile)
	if err != nil {
		common.Failure(c, err)
		return
	}
	err = writer.WriteField("question", question)
	if err != nil {
		common.Failure(c, err)
		return
	}
	_ = imageFile.Close()
	_ = writer.Close()

	req, err = http.NewRequest("POST", vqaimage2, buf)
	if err != nil {
		common.Failure(c, err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Auth-Token", t)

	response, err = client.Do(req)
	if err != nil {
		common.Failure(c, err)
		return
	}

	bys, err = ioutil.ReadAll(response.Body)
	if err != nil {
		common.Failure(c, err)
		return
	}

	if response.StatusCode >= 300 {
		common.Failure(c, errors.New(string(bys)))
		return
	}

	var ret demo.VqaRes
	if err = json.Unmarshal(bys, &ret); err != nil {
		common.Failure(c, err)
		return
	}

	if err = ret.Valication(); err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, ret)
}

func (h *Handle) Vqa(c *gin.Context) {
	v := c.Param("v")
	if v == "2" {
		h.vqa2(c)
	} else {
		h.vqa4(c, v)
	}
}

func (h *Handle) vqa4(c *gin.Context, v string) {
	var req demo.VqaReq
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		common.QueryFailure(c, err)
		return
	}
	t := h.gettoken()
	var (
		request *http.Request
		resp    *http.Response
		bys     []byte
		url     string
		err     error
	)
	url = vqaimage4
	if v == "3" {
		url = vqaimage3
	}

	request, err = http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf(`{"image_path":"%s","question":"%s"}`, req.Image, req.Question)))
	if err != nil {
		common.Failure(c, err)
		return
	}

	request.Header.Set("X-Auth-Token", t)
	request.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(request)
	if err != nil {
		common.Failure(c, err)
		return
	}

	bys, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		common.Failure(c, err)
		return
	}

	if resp.StatusCode >= 300 {
		common.Failure(c, errors.New(string(bys)))
		return
	}

	var ret demo.VqaRes
	if err = json.Unmarshal(bys, &ret); err != nil {
		common.Failure(c, err)
		return
	}

	if err = ret.Valication(); err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, ret)

}
