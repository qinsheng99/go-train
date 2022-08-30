package model

type Boy struct {
	Id          int64  `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Information string `json:"information" gorm:"column:information"`
}

func (Boy) TableName() string {
	return "boy"
}
