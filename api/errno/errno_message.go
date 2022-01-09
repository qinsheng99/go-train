package errno

//import "errno"

var (
	OK        = NewError(0, "success")
	CookieErr = NewError(900, "未登录")
	DataErr   = NewError(122, "数据参数不正确，请勿非法操作")
)
