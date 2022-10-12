package model

import (
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
)

type Boy struct {
	Id           int64         `json:"id" gorm:"column:id"`
	Name         string        `json:"name" gorm:"column:name"`
	Informations pgtype.JSONB  `gorm:"column:information" json:"information"`
	Arr          pq.Int64Array `gorm:"column:arr;type:integer[]" json:"arr"`
	//Information  Infor         `json:"information" gorm:"-"`
}

type OldBoy struct {
	Id   int64         `json:"id" gorm:"column:id"`
	Name string        `json:"name" gorm:"column:name"`
	Arr  pq.Int64Array `gorm:"column:arr;type:integer[]" json:"arr"`
	//Information  Infor         `json:"information" gorm:"-"`
}

func (b *Boy) TableName() string {
	return "boy"
}

func (b *OldBoy) TableName() string {
	return "oldboy"
}

type Infor struct {
	Age     int    `json:"age"`
	Address string `json:"address"`
}
