package entity

type Response struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	NowTime int64       `json:"nowTime"`
	Data    interface{} `json:"data"`
}

/**
"code": 0,
    "msg": "success",
    "nowTime": 1640242745,
    "data": "151****1596"

*/
