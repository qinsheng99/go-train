package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qinsheng99/goWeb/library/config"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BundleDb struct {
	Db *gorm.DB
}

const CONNMAXLIFTIME = 900

func GetBundleDb(cfg *config.MysqlConfig) (*BundleDb, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPwd, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := gorm.Open(gormmysql.New(gormmysql.Config{
		DSN:                       dsn,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(CONNMAXLIFTIME)
	sqlDB.SetMaxOpenConns(cfg.DbMaxConn)
	sqlDB.SetMaxIdleConns(cfg.DbMaxidle)
	return &BundleDb{
		Db: db,
	}, nil
}
