package httprequest

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/qinsheng99/goWeb/library/logger"
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
	var req *http.Request
	var resp *http.Response
	err = try.Do(func(attempt int) (retry bool, err error) {
		req, err = http.NewRequest(method, url, getbody(bytesData))
		if err != nil {
			logger.SLog.Errorf("reqURL:%s ;http new request err: %v", url, err)
			return attempt < 3, err
		}

		req.Header.Set("Content-Type", "application/json")
		for key, item := range headers {
			req.Header.Set(key, item)
		}

		resp, err = client.Do(req)
		if err != nil || resp == nil {
			return attempt < 3, err
		}
		defer resp.Body.Close()
		resByte, err = ioutil.ReadAll(resp.Body)

		if resp.StatusCode > http.StatusMultipleChoices || resp.Body == nil {
			logger.SLog.Error(fmt.Sprintf("statusCode is %d ,data : %s", resp.StatusCode, string(resByte)))
			return attempt < 3, errors.New(fmt.Sprintf("statusCode is %d ,data : %s", resp.StatusCode, string(resByte)))
		}

		return attempt < 3, err
	})
	return
}

func Get(url string, bytesData interface{}, headers map[string]string, u url.Values) ([]byte, error) {
	return mainRequest(geturl(url, u), "GET", bytesData, headers)
}

func Post(url string, bytesData interface{}, headers map[string]string, u url.Values) ([]byte, error) {
	return mainRequest(geturl(url, u), "POST", bytesData, headers)
}

// NoTryRequest 所有公用的http请求无重试
func NoTryRequest(url, method string, bytesData interface{}, headers map[string]string) (resByte []byte, err error) {
	req, err := http.NewRequest(method, url, getbody(bytesData))
	if err != nil {
		logger.SLog.Errorf("reqURL:%s ;http new request err: %v", url, err)
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

	if resp.StatusCode > http.StatusMultipleChoices || resp.Body == nil {
		return nil, errors.New("响应状态码有误")
	}
	resByte, err = ioutil.ReadAll(resp.Body)
	return
}

func getbody(bytesData interface{}) io.Reader {
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
	return body
}

func geturl(u string, values url.Values) string {
	path, err := url.Parse(u)
	if err != nil {
		return u
	}

	if len(values) > 0 {
		q := path.Query()

		for s, value := range values {
			for _, v := range value {
				q.Add(s, v)
			}
		}
		path.RawQuery = q.Encode()
	}
	return path.String()
}
