package tools

// import (
// 	"github.com/qinsheng99/goWeb/err"
// 	"net/http"
//
// 	"github.com/gin-gonic/gin"
// 	"gitlab.weipaitang.com/tech-lab/dionysus/pkg"
// )
//
// var ParamsErr = ErrorsNew(err.DataStatus)
// var ThirdErr = ErrorsNew(err.ThirdError)
//
// // 系统异常
// var ServerErr = ErrorsNew(err.Error)
//
// // 临时定义，解决项目中自动生成 code,msg
// type Error struct {
// 	Id     string `json:"id"`
// 	Code   int32  `json:"code"`
// 	Detail string `json:"detail"`
// 	Status string `json:"status"`
// }
//
// type BadInfo struct {
// 	Code int         `json:"code"`
// 	Data interface{} `json:"data"`
// 	Err  error       `json:"err"`
// }
//
// funcTest (e *Error) Error() string {
// 	return e.Detail
// }
//
// funcTest New(id, detail string, code int32) error {
// 	return &Error{
// 		Id:     id,
// 		Code:   code,
// 		Detail: detail,
// 		Status: http.StatusText(int(code)),
// 	}
// }
//
// funcTest ParamError(c *gin.Context) {
// 	c.Render(200, pkg.json{Data: BadReturn(ParamsErr, nil)})
// }
// funcTest HandleError(c *gin.Context, err error) {
// 	c.Render(200, pkg.JSON{Data: HandleBadReturn(err, nil)})
// }
//
// funcTest HandleSuccess(c *gin.Context, data interface{}, options ...RespOption) {
// 	c.Render(200, pkg.JSON{Data: SuccessReturn(nil, data, options...)})
// }
//
// funcTest HandleProxyResp(c *gin.Context, data interface{}, options ...RespOption) {
// 	c.Render(200, pkg.JSON{Data: data})
// }
//
// funcTest BadReturn(err error, data interface{}) map[string]interface{} {
// 	var info = make(map[string]interface{})
// 	info["code"] = errorBase.DataStatus
// 	info["err"] = err.Error()
// 	info["data"] = data
// 	return info
// }
//
// funcTest HandleBadReturn(err error, data interface{}) map[string]interface{} {
// 	var info = make(map[string]interface{})
// 	info["code"] = err.Error
// 	info["err"] = err.Error()
// 	info["data"] = data
// 	return info
// }
//
// type RespOption funcTest(map[string]interface{})
//
// type ErrOption funcTest(map[string]interface{})
//
// funcTest OptionCode(code int64) RespOption {
// 	return funcTest(m map[string]interface{}) {
// 		m["code"] = code
// 	}
// }
// funcTest OptionErrMsg(errMsg string) RespOption {
// 	return funcTest(m map[string]interface{}) {
// 		m["err"] = errMsg
// 	}
// }
//
// funcTest SuccessReturn(err error, data interface{}, options ...RespOption) map[string]interface{} {
// 	var info = make(map[string]interface{})
// 	info["code"] = errorBase.Success
// 	info["err"] = err
// 	info["data"] = data
// 	for _, option := range options {
// 		option(info)
// 	}
// 	return info
// }
//
// funcTest NewError(code int32, text string) error {
// 	return New(constant.ProjectName, text, code)
// }
//
// funcTest ErrorsNew(code int) error {
// 	if code < 0 {
// 		return NewError(int32(code), errorBase.ErrorMsg(code))
// 	}
// 	return New(constant.ProjectName, errorBase.ErrorMsg(code), int32(code))
// }
//
// funcTest ErrorsNewMsg(code int, msg string) error {
// 	if code < 0 {
// 		return NewError(int32(code), msg)
// 	}
// 	return New(constant.ProjectName, msg, int32(code))
// }
//
// funcTest NotifySuccess(c *gin.Context, data string) {
// 	c.Render(http.StatusOK, pkg.String{Format: data})
// 	c.Abort()
// }
//
// funcTest NotifyFail(c *gin.Context, data string) {
// 	c.Render(http.StatusOK, pkg.String{Format: data})
// 	c.Abort()
// }
//
// funcTest NotifyError(text string) error {
// 	return New(constant.ProjectName, text, http.StatusServiceUnavailable)
// }
//
// type Response struct {
// 	Code    int32                  `json:"code"`
// 	Msg     string                 `json:"msg"`
// 	NowTime int64                  `json:"nowTime,omitempty"`
// 	Data    map[string]interface{} `json:"data,omitempty"`
// }
