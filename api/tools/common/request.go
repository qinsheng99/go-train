package common

import (
	"net/http"
	"strconv"
	"time"

	"github.com/qinsheng99/go-train/api/errno"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}, options ...RespOption) {
	c.JSON(http.StatusOK, SuccessReturn(data, c, options...))
}

func Failure(c *gin.Context, err error) {
	c.JSON(http.StatusOK, HandleBadReturn(err, nil))
}

func QueryFailure(c *gin.Context, err error) {
	c.JSON(http.StatusOK, QueryHandleBadReturn(err, nil))
}

type RespOption func(m map[string]interface{})

func SuccessReturn(data interface{}, c *gin.Context, options ...RespOption) map[string]interface{} {
	var info = make(map[string]interface{})
	start := c.GetTime("start")
	info["code"] = errno.OK.Code
	info["msg"] = errno.ErrorMsg(errno.OK.Code)
	info["nowTime"] = time.Now().Unix()
	info["costTime"] = strconv.Itoa(int(time.Now().Sub(start).Milliseconds())) + "ms"
	info["data"] = data
	for _, option := range options {
		option(info)
	}
	return info
}

func HandleBadReturn(err error, data interface{}) map[string]interface{} {
	var info = make(map[string]interface{})
	info["code"] = errno.Error
	info["msg"] = err.Error()
	info["nowTime"] = time.Now().Unix()
	info["data"] = data
	return info
}

func QueryHandleBadReturn(err error, data interface{}) map[string]interface{} {
	var info = make(map[string]interface{})
	info["code"] = errno.DataErr.Code
	info["msg"] = errno.ErrorMsg(errno.DataErr.Code)
	info["nowTime"] = time.Now().Unix()
	info["err"] = err.Error()
	info["data"] = data
	return info
}
