package httprequest

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/qinsheng99/goWeb/library/try"
)

var (
	client      *http.Client
	clientNoTry *http.Client
)

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        250,
			MaxIdleConnsPerHost: 250,
			IdleConnTimeout:     time.Duration(120) * time.Second,
			DisableKeepAlives:   false,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
	}
	clientNoTry = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        250,
			MaxIdleConnsPerHost: 250,
			IdleConnTimeout:     time.Duration(120) * time.Second,
			DisableKeepAlives:   false,
		},
	}
}

// MainRequest 所有公用的http请求
func mainRequest(url, method string, bytesData interface{}, headers map[string]string) (resByte []byte, err error) {
	var body = io.Reader(nil)
	switch t := bytesData.(type) {
	case []byte:
		body = bytes.NewReader(t)
	case string:
		body = strings.NewReader(t)
	case *strings.Reader:
		body = t
	case *bytes.Buffer:
		body = t
	default:
		body = nil
	}
	err = try.Do(func(attempt int) (retry bool, err error) {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			//logger.Warnf("reqURL:%s ;http new request err: %v", url, err)
			return attempt < 3, err
		}
		// 本地测试使用
		// req.Header.Set("Cookie", "wpt_env_num=test-13")
		req.Header.Set("Content-Type", "application/json")
		for key, item := range headers {
			req.Header.Set(key, item)
		}

		resp, err := client.Do(req)
		if err != nil || resp == nil {
			//logger.Infof("请求reqURL:%s header:%+v 参数:%s", url, headers, string(bytesData))
			return attempt < 3, err
		}
		defer resp.Body.Close()
		resByte, err = ioutil.ReadAll(resp.Body)

		if resp.StatusCode > http.StatusMultipleChoices || resp.Body == nil {
			return attempt < 3, errors.New("响应状态码有误")
		}

		return attempt < 3, err
	})
	return
}

func Get(url string, bytesData interface{}, headers map[string]string) ([]byte, error) {
	return mainRequest(url, "GET", bytesData, headers)
}

func Post(url string, bytesData interface{}, headers map[string]string) ([]byte, error) {
	return mainRequest(url, "POST", bytesData, headers)
}

// NoTryRequest 所有公用的http请求无重试
func NoTryRequest(url, method string, bytesData interface{}, headers map[string]string) (resByte []byte, err error) {
	var body = io.Reader(nil)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		//logger.Warnf("reqURL:%s ;http new request err: %v", url, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	for key, item := range headers {
		req.Header.Set(key, item)
	}
	resp, err := clientNoTry.Do(req)
	if err != nil || resp == nil {
		//logger.Warnf("reqURL:%s ;client Do err: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK || resp.Body == nil {
		//logger.Infof("reqURL:%s ;resp status code is not 200 ,resp: %v", url, resp)
	}
	resByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		//logger.Errorf("error:%v", err)
	}
	return
}
