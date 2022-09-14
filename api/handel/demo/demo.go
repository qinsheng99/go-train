package demo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/goWeb/api/tools/common"
)

type Handle struct {
}

func NewDemo() *Handle {
	return &Handle{}
}

type res struct {
	Result r `json:"inference_result"`
}

type r struct {
	Instances instances `json:"instances"`
}

type instances struct {
	Image []string `json:"image"`
}

type AuthToken struct {
}

var client *http.Client

func init() {
	timeout := time.Duration(10 * time.Second)
	client = &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func (h *Handle) File(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		common.Failure(c, err)
		return
	}
	path := "../images/" + f.Filename
	err = c.SaveUploadedFile(f, path)
	if err != nil {
		common.Failure(c, err)
		return
	}
	resp, err := h.send(path)
	if err != nil {
		common.Failure(c, err)
		return
	}
	//defer os.Remove(path)
	all, _ := ioutil.ReadAll(resp.Body)
	var re res
	json.Unmarshal(all, &re)
	fmt.Println(string(all))
	common.Success(c, re)
}

func token() (string, error) {
	body := fmt.Sprintf(`{"auth": {"identity": {"methods": ["password"],"password": {"user": {"name": %v,"password": "%v","domain": {"name": "%v"}}}},"scope": {"project": {"name": "cn-central-221"}}}}`,
		"wuhanjisuan191", "986%#hwAA1", "wuhanjisuan191")
	post, err := http.Post("https://iam-pub.cn-central-221.ovaijisuan.com/v3/auth/tokens", "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	return post.Header.Get("x-subject-token"), nil
}

func (h *Handle) send(path string) (*http.Response, error) {
	t, err := token()
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	file, _ := writer.CreateFormFile("file", "xiaohu.png")
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	io.Copy(file, imageFile)
	imageFile.Close()
	writer.Close()
	req, err := http.NewRequest("POST", "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/ef220239-dfeb-4400-96b4-5fe0d4b35735/infer/image", buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Auth-Token", t)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
