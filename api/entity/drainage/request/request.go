package request

type DrainageRequest struct {
	Type int `form:"type" binding:"required,gt=3" json:"type" xml:"type"`
	Id   int `form:"id" binding:"required" json:"id" xml:"id"`
}
