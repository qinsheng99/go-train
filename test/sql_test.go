package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qinsheng99/go-train/internal/model"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgreSqlCreate(t *testing.T) {
	u, _ := uuid.NewRandom()
	if err := d.Exec(d.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&model.Boy{
			Name: "xiaohu13",
			Json: postgres.Jsonb{RawMessage: []byte(`{"age": 23, "repo": "rng"}`)},
			Arr:  []int64{13, 26, 39},
			UUid: u,
		})
	})).Error; err != nil {
		t.Fatal(err)
	}
}

func TestPostgreSqlUpdate(t *testing.T) {
	b := &model.Boy{
		Informations: postgres.Jsonb{RawMessage: []byte(`{"age": 233, "repo": "rngay"}`)},
	}
	t.Log(d.Exec(d.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(b).Where("id = 12").Updates(b)
	})).Error)
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
