//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"gin/api/handel/ceshi"
	"gin/api/handel/mysql"
	"gin/api/handel/redis"
	"gin/internal/dao/persistence"
	"gin/internal/dao/persistence/customer"
	ceshi2 "gin/internal/service/ceshi"
	"gin/internal/service/drainage"
	ServiceMysql "gin/internal/service/mysql"
	"gin/library/db"
	"gin/library/elasticsearch"
	"gin/library/redisClient"
)

func Init(bundleDb *db.BundleDb, es *elasticsearch.ES, r *redisClient.Redis) (*Entry, error) {
	NewEsDao := persistence.NewEsDao(es)
	NewCustomerDao := customer.NewCustomerDao(bundleDb, NewEsDao)
	Drainage := drainage.NewDS(NewCustomerDao)
	NewHandlerDao := ceshi.NewHandler(NewCustomerDao, NewEsDao, Drainage, bundleDb, r)
	NewCeshi := ceshi2.NewCeshi()
	NewRedis := redisClient.NewRedisStruct(r)
	NewH := redis.NewH(NewCeshi, NewRedis)
	NewMysqlImp := ServiceMysql.NewMysqlService(bundleDb)
	NewMysql := mysql.NewMysql(NewMysqlImp)

	e := &Entry{
		NewHandler: NewHandlerDao,
		NewH:       NewH,
		NewMysql:   NewMysql,
	}

	return e, nil
}
