package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BundleDb struct {
	Db *gorm.DB
}

const (
	BundleDbURI = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s"
)

func GetBundleDb() (*BundleDb, error) {
	// db, err := gorm.Open(mysql.Open(fmt.Sprintf(BundleDbURI, "dev", "5f031c352d1a4ebd", "106.14.210.131", "3306", "dev")))
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(BundleDbURI, "root", "12345678", "localhost", "3306", "ceshi")))
	if err != nil {
		fmt.Printf("mysql connect error: %v", err)
		return nil, err
	}
	if db.Error != nil {
		fmt.Printf("mysql connect error: %v", err)
		return nil, err
	}
	return &BundleDb{
		Db: db,
	}, nil
}
