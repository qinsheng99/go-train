package test

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qinsheng99/goWeb/internal/model"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgreSql(t *testing.T) {
	if err := d.Create(&model.Boy{
		Name:         "xiaohu13",
		Informations: postgres.Jsonb{RawMessage: []byte(`{"age": 23, "repo": "rng"}`)},
		Arr:          []int64{13, 26, 39},
	}).Error; err != nil {
		t.Fatal(err)
	}
}

var d *gorm.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", "postgres", "root", "postgres", 5432)
	db, err := gorm.Open(p.New(p.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return
	}
	d = db
	m.Run()
}
