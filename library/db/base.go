package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qinsheng99/go-train/library/config"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BundleDb struct {
	Db *gorm.DB
}

var mysqlDb *gorm.DB

var postgresqlDb *gorm.DB

const CONNMAXLIFTIME = 900

func GetMysql(cfg *config.MysqlConfig) (*BundleDb, error) {
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
	mysqlDb = db
	return &BundleDb{
		Db: db,
	}, nil
}

func GetPostgresql(cfg *config.PostgresqlConfig) error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", cfg.DbHost, cfg.DbUser, cfg.DbPwd, cfg.DbName, cfg.DbPort)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxLifetime(CONNMAXLIFTIME)
	sqlDB.SetMaxOpenConns(cfg.DbMaxConn)
	sqlDB.SetMaxIdleConns(cfg.DbMaxidle)
	postgresqlDb = db
	return nil
}

func GetMysqlDb() *gorm.DB {
	return mysqlDb
}

func GetPostgresqlDb() *gorm.DB {
	return postgresqlDb
}
