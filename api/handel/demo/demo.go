package demo

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/library/redisClient"
)

type Handle struct {
	redis redisClient.RedisInterface
}

func NewDemo(red redisClient.RedisInterface) *Handle {
	return &Handle{redis: red}
}

var (
	oneimage = []string{
		"https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/43a04dbe-c94e-41ba-a0e5-9da34efa8ff3/text2image",
		//"https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/b3433d2a-6320-4171-a687-bce38e3a9eca/text2image",
	}
	threeimage = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/a55e424e-b4fd-403b-97f9-d406be420f84/text2image"
)

var mmap sync.Map

func init() {
	mmap.Store("https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/43a04dbe-c94e-41ba-a0e5-9da34efa8ff3/text2image", true)
	mmap.Store("https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/b3433d2a-6320-4171-a687-bce38e3a9eca/text2image", true)
}

var lockthree sync.Mutex
var lockone sync.Mutex

type textImageRes struct {
	Status int         `json:"status"`
	Output interface{} `json:"output_image_url"`
	Msg    string      `json:"msg"`
}

func (t *textImageRes) validation() error {
	if t.Status == -1 {
		return errors.New(t.Msg)
	}
	return nil
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

type testImage struct {
	Input string `json:"input_text" binding:"required"`
	User  string `json:"user_name" binding:"required"`
	Flag  bool   `json:"flag"`
}

var client *http.Client

func init() {
	timeout := time.Duration(2 * time.Minute)
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
	if err = json.Unmarshal(all, &re); err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, re)
}

func token() (string, error) {
	body := fmt.Sprintf(`{"auth": {"identity": {"methods": ["password"],"password": {"user": {"name": "%v","password": "%v","domain": {"name": "%v"}}}},"scope": {"project": {"name": "cn-central-221"}}}}`,
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

func (h *Handle) textimage(url, token string, req testImage) ([]byte, error) {
	var data []byte
	bys := []byte(fmt.Sprintf(`{"input_text":"%v","user_name":"%v"}`, req.Input, req.User))

	re, err := http.NewRequest("POST", url, bytes.NewReader(bys))
	if err != nil {
		return nil, err
	}
	re.Header.Set("Content-Type", "application/json")
	re.Header.Set("X-Auth-Token", token)

	response, err := client.Do(re)
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 200 {
		return nil, errors.New(string(data))
	}
	return data, nil
}

func (h *Handle) geturl() (string, error) {
	var url string
	for _, b := range oneimage {
		if f, ok := mmap.Load(b); ok && f.(bool) {
			url = b
			mmap.Store(b, false)
			goto LOOP
		}
	}
	return "", errors.New("当前节点没有可用的,请稍后再试")
LOOP:
	return url, nil
}

func (h *Handle) TestImage(c *gin.Context) {
	var req testImage
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.QueryFailure(c, err)
		return
	}

	t, err := h.redis.Get(context.Background(), "modelarts-token")
	if len(t) == 0 || err != nil {
		t, err = token()
		if err != nil {
			common.Failure(c, err)
			return
		}
		_, _ = h.redis.Set(context.Background(), "modelarts-token", t, time.Hour*24)
	}

	var url string
	var data []byte
	if req.Flag {
		lockthree.Lock()
		data, err = h.textimage(threeimage, t, req)
		lockthree.Unlock()
		if err != nil {
			common.Failure(c, err)
			return
		}
	} else {
		url, err = h.geturl()
		if err != nil {
			common.Failure(c, err)
			return
		}
		data, err = h.textimage(url, t, req)
		mmap.Store(url, true)
		if err != nil {
			common.Failure(c, err)
			return
		}
	}

	var result textImageRes
	//log.Println(string(data))
	if err = json.Unmarshal(data, &result); err != nil {
		common.Failure(c, err)
		return
	}

	if err = result.validation(); err != nil {
		common.Failure(c, err)
		return
	}
	common.Success(c, result)
}
