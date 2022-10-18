package demo

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	httprequest "github.com/qinsheng99/goWeb/library/request"
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
		"https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/b3433d2a-6320-4171-a687-bce38e3a9eca/text2image",
	}
	threeimage = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/a55e424e-b4fd-403b-97f9-d406be420f84/text2image"

	vqaimage2 = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/b6c5cf73-de6a-49ed-ac40-5e943903e010/v2/infer/vqa"

	vqaimage4 = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/b6c5cf73-de6a-49ed-ac40-5e943903e010/v4/infer/vqa"
	vqaimage3 = "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/b6c5cf73-de6a-49ed-ac40-5e943903e010/v3/infer/vqa"
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
	client = &http.Client{
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
	if f.Size >= 200000 {
		common.Failure(c, errors.New("the file size cannot be larger than 200kb"))
		return
	}
	path := "../images/" + f.Filename
	err = c.SaveUploadedFile(f, path)
	if err != nil {
		common.Failure(c, err)
		return
	}
	all, err := h.send(path)
	if err != nil {
		common.Failure(c, err)
		return
	}
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

func (h *Handle) send(path string) ([]byte, error) {
	var data []byte
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
	head := map[string]string{
		"X-Auth-Token": t,
		"Content-Type": writer.FormDataContentType(),
	}
	u := "https://a2f051d4cabf45f885d7b0108edc9b9c.infer.ovaijisuan.com/v1/infers/ef220239-dfeb-4400-96b4-5fe0d4b35735/infer/image"
	data, err = httprequest.Post(u, buf, head, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *Handle) textimage(url, token string, req testImage) ([]byte, error) {
	var data []byte
	var err error

	head := map[string]string{
		"X-Auth-Token": token,
	}

	data, err = httprequest.Post(url, []byte(fmt.Sprintf(`{"input_text":"%v","user_name":"%v"}`, req.Input, req.User)), head, nil)
	if err != nil {
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
