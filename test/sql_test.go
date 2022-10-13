package test

import (
	"fmt"
	"testing"

	"github.com/qinsheng99/goWeb/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgreSql(t *testing.T) {
	if err := d.Exec(d.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = 10").Updates(&model.Boy{
			Name:         "xiaohu10",
			Informations: model.Jsonb{RawMessage: []byte(`{"age": 23, "repo": "rng"}`)},
			Arr:          []int64{10, 20, 30},
		})
	})).Error; err != nil {
		t.Fatal(err)
	}
}

var d *gorm.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", "postgres", "root", "postgres", 5432)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return
	}
	d = db
	m.Run()
}
