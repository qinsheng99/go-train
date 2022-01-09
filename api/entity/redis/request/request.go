package request

type Hset struct {
	Name  string `json:"name" binding:"required"`
	Field string `json:"field"`
	Data  int    `json:"data"`
	Bo    bool   `json:"bo"`
}

type RandMember struct {
	Count int64 `form:"count"`
}
