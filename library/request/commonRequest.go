package httprequest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/qinsheng99/goWeb/library/try"
)

var (
	client      *http.Client
	clientNoTry *http.Client
)

func init() {
	client = &http.Client{
		Timeout: 2 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        250,
			MaxIdleConnsPerHost: 250,
			IdleConnTimeout:     time.Duration(120) * time.Second,
			DisableKeepAlives:   false,
		},
	}
	clientNoTry = &http.Client{Timeout: 8 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        250,
			MaxIdleConnsPerHost: 250,
			IdleConnTimeout:     time.Duration(120) * time.Second,
			DisableKeepAlives:   false,
		},
	}
}

// MainRequest 所有公用的http请求
func mainRequest(url, method string, bytesData []byte, headers map[string]string) (resByte []byte, err error) {
	err = try.Do(func(attempt int) (retry bool, err error) {
		req, err := http.NewRequest(method, url, bytes.NewReader(bytesData))
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

		if resp.StatusCode != http.StatusOK || resp.Body == nil {
			//logger.Infof("请求reqURL:%s header:%+v 参数:%s 返回:%s", url, headers, string(bytesData), string(resByte))
			return attempt < 3, errors.New("响应状态码不是200")
		}

		return attempt < 3, err
	})

	if err != nil {
		//logger.Errorf("error:%v", err)
	}
	return
}

func Get(url string, bytesData []byte, headers map[string]string) ([]byte, error) {
	return mainRequest(url, "GET", bytesData, headers)
}

func Post(url string, bytesData []byte, headers map[string]string) ([]byte, error) {
	return mainRequest(url, "POST", bytesData, headers)
}

// NoTryRequest 所有公用的http请求无重试
func NoTryRequest(url, method string, bytesData []byte, headers map[string]string) (resByte []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(bytesData))
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
