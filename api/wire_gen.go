//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"github.com/qinsheng99/goWeb/api/handel/ceshi"
	"github.com/qinsheng99/goWeb/api/handel/mysql"
	"github.com/qinsheng99/goWeb/api/handel/redis"
	sortHandler "github.com/qinsheng99/goWeb/api/handel/sort"
	"github.com/qinsheng99/goWeb/internal/dao/persistence"
	"github.com/qinsheng99/goWeb/internal/dao/persistence/customer"
	ceshi2 "github.com/qinsheng99/goWeb/internal/service/ceshi"
	"github.com/qinsheng99/goWeb/internal/service/drainage"
	ServiceMysql "github.com/qinsheng99/goWeb/internal/service/mysql"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/elasticsearch"
	"github.com/qinsheng99/goWeb/library/redisClient"
)

func Init(bundleDb *db.BundleDb, es *elasticsearch.ES, r *redisClient.Redis) (*Entry, error) {
	NewEsDao := persistence.NewEsDao(es)
	NewCustomerDao := customer.NewCustomerDao(bundleDb, NewEsDao)
	Drainage := drainage.NewDS(NewCustomerDao)
	NewHandlerDao := ceshi.NewHandler(NewCustomerDao, NewEsDao, Drainage)
	NewCeshi := ceshi2.NewCeshi()
	NewRedis := redisClient.NewRedisStruct(r)
	NewH := redis.NewH(NewCeshi, NewRedis)
	NewMysqlImp := ServiceMysql.NewMysqlService(bundleDb)
	NewMysql := mysql.NewMysql(NewMysqlImp)
	NewSort := sortHandler.NewSort()

	e := &Entry{
		NewHandler: NewHandlerDao,
		NewH:       NewH,
		NewMysql:   NewMysql,
		NewSort: NewSort,
	}

	return e, nil
}
