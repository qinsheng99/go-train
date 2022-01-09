package model

type Record struct {
	Id         int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Uri        string `json:"uri" gorm:"column:uri"`
	IsDelete   int    `json:"is_delete" gorm:"column:is_delete"`
	BackName   string `json:"back_name" gorm:"column:back_name"`
	CreateTime string `json:"create_time" gorm:"column:create_time"`
	DeleteTime string `json:"delete_time" gorm:"column:delete_time"`
}

func (Record) TableName() string {
	return "ceshi"
}
