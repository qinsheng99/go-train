package request

type CeShiRequest struct {
	Name       string `json:"name" form:"name" binding:"-"`
	Ids        []int  `json:"id" form:"id" binding:"-"`
	Gender     int    `json:"gender" form:"gender" binding:"-"`
	Page       int    `json:"page" gorm:"page" binding:"-"`
	PageSize   int    `json:"pageSize" form:"pageSize" binding:"-"`
	UserId     string `json:"user_id" binding:"-" form:"user_id"`
	StaffId    int    `json:"staff_id" binding:"-" form:"staff_id"`
	AddTime    uint   `json:"add_time" binding:"-" form:"add_time"`
	CompanyId  int    `json:"company_id" binding:"-" form:"company_id"`
	AddChannel []int  `json:"add_channel" binding:"-" form:"add_channel"`
}

type CeShiGetRequest struct {
	Name       string `json:"name" form:"name" binding:"-"`
	Ids        []int  `json:"ids" form:"ids" binding:"-"`
	Gender     int    `json:"gender" form:"gender" binding:"-"`
	Page       int    `json:"page" gorm:"page" binding:"-"`
	PageSize   int    `json:"pageSize" form:"pageSize" binding:"-"`
	UserId     string `json:"user_id" binding:"-" form:"user_id"`
	StaffId    int    `json:"staff_id" binding:"-" form:"staff_id"`
	AddTime    uint   `json:"add_time" binding:"-" form:"add_time"`
	CompanyId  int    `json:"company_id" binding:"-" form:"company_id"`
	AddChannel []int  `json:"add_channel" binding:"-" form:"add_channel"`
	Tags       []int  `json:"tags" form:"tags"`
	Id         int    `json:"id" form:"id"`
	Option     string `json:"option" form:"option"`
	LossState  int    `json:"loss_state" form:"lossState"`
}
