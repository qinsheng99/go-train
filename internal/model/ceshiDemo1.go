package model

type CeshiDemo1 struct {
	Id     int    `gorm:"AUTO_INCREMENT;column:id;type:INT;primary_key" json:"id"`
	Uri    string `gorm:"column:uri" json:"uri"`
	Name   string `gorm:"column:name" json:"name"`
	Status int    `json:"status" gorm:"column:status"`
}

func (CeshiDemo1) TableName() string {
	return "ceshi_demo1"
}

type CeshiJoin struct {
	Ceshi
	Name     string `json:"name"`
	NameName string `json:"name_name"`
}
