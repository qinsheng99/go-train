package model

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"github.com/qinsheng99/go-train/library/db"
	"gorm.io/gorm"
)

type Boy struct {
	Id            int64          `json:"id,omitempty" gorm:"column:id;type:int8"`
	Name          string         `json:"name,omitempty" gorm:"column:name;type:varchar(30)"`
	Informations  postgres.Jsonb `gorm:"column:information;type:jsonb;default:'{}'" json:"-"`
	Informationss string         `gorm:"-" json:"information,omitempty"`
	Json          postgres.Jsonb `gorm:"column:jsondata;type:json;default:'{}'" json:"-"`
	Jsons         string         `gorm:"-" json:"json,omitempty"`
	Arr           pq.Int64Array  `gorm:"column:arr;type:integer[]" json:"arr,omitempty"`
	UUid          uuid.UUID      `gorm:"column:uuid;type:uuid" json:"uuid,omitempty"`
}

type BoyArr struct {
	Boy
	Arrone int64 `json:"arrone" gorm:"column:arrone"`
}

func (b *Boy) TableName() string {
	return "boy"
}

type Jsonb struct {
	json.RawMessage
}

func (b *Boy) Insert() (*Boy, error) {
	cli := db.GetPostgresqlDb()
	if err := cli.Exec(cli.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(b)
	})).Error; err != nil {
		return nil, err
	}
	cli.Last(b)
	return b, nil
}

func (b *Boy) Update() (err error) {
	cli := db.GetPostgresqlDb()
	err = cli.Exec(cli.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", b.Id).UpdateColumns(
			map[string]interface{}{
				"information": b.Informations,
			},
		)
	})).Error

	return
}
