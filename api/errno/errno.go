package errno

/*
*错误码
 */
const (
	Error         = 200
	Success       = 0
	NotLoginError = 900
	MissingData   = 400001
	DataStatus    = 122
	ParamIllegal  = 400004
	RedisError    = 400005
	ParamError    = 400006
	ThirdError    = 500005

	ApplyTokenErr = 42001
)

var errorMsg = map[int]string{
	Error:         "系统异常",
	Success:       "success",
	NotLoginError: "未登录",
	MissingData:   "数据缺失",
	DataStatus:    "数据参数不正确，请勿非法操作",
	ParamIllegal:  "参数传入不合法:[%s]",
	RedisError:    "redis连接操作失败",
	ParamError:    "缺失参数不能",
	ThirdError:    "请求第三方服务错误",
	ApplyTokenErr: "invalid credential, access_token is invalid or not latest",
}

type Err struct {
	Code int
	Msg  string
}

func ErrorMsg(code int) string {
	return errorMsg[code]
}

func (e Err) Error() string {
	return e.Msg
}

func NewCode(a int) *Err {
	return &Err{
		Code: a,
	}
}

type NMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(i int, s string) *NMsg {
	return &NMsg{
		Code: i,
		Msg:  s,
	}
}
