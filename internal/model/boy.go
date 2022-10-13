package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type Boy struct {
	Id           int64         `json:"id" gorm:"column:id;type:int8"`
	Name         string        `json:"name" gorm:"column:name;type:varchar(30)"`
	Informations Jsonb         `gorm:"column:information;type:jsonb json:"information"`
	Arr          pq.Int64Array `gorm:"column:arr;type:integer[]" json:"arr"`
	//Information  Infor         `json:"information" gorm:"-"`
}
type Jsonb struct {
	json.RawMessage
}

// Value get value of Jsonb
func (j Jsonb) Value() (driver.Value, error) {
	if len(j.RawMessage) == 0 {
		return nil, nil
	}
	return j.MarshalJSON()
}

// Scan scan value into Jsonb
func (j *Jsonb) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, j)
}

type OldBoy struct {
	Id     int64         `json:"id" gorm:"column:id"`
	Name   string        `json:"name" gorm:"column:name"`
	Arr    pq.Int64Array `gorm:"column:arr;type:integer[]" json:"arr"`
	Create time.Time     `gorm:"column:create_time" json:"create"`
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
