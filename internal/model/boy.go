package model

type Boy struct {
	Id           int64  `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`
	Informations []byte `gorm:"column:information" json:"-"`
	Information  Infor  `json:"information" gorm:"-"`
}

func (Boy) TableName() string {
	return "boy"
}

type Infor struct {
	Age     int    `json:"age"`
	Address string `json:"address"`
}
