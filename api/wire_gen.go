//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"github.com/qinsheng99/goWeb/api/handel/ceshi"
	"github.com/qinsheng99/goWeb/api/handel/demo"
	esHandle "github.com/qinsheng99/goWeb/api/handel/es"
	"github.com/qinsheng99/goWeb/api/handel/mongo"
	"github.com/qinsheng99/goWeb/api/handel/mysql"
	"github.com/qinsheng99/goWeb/api/handel/redis"
	sortHandler "github.com/qinsheng99/goWeb/api/handel/sort"
	"github.com/qinsheng99/goWeb/internal/dao/persistence"
	"github.com/qinsheng99/goWeb/internal/dao/persistence/boy"
	"github.com/qinsheng99/goWeb/internal/dao/persistence/customer"
	ceshi2 "github.com/qinsheng99/goWeb/internal/service/ceshi"
	"github.com/qinsheng99/goWeb/internal/service/drainage"
	ServiceMysql "github.com/qinsheng99/goWeb/internal/service/mysql"
	servicePostgresql "github.com/qinsheng99/goWeb/internal/service/postgresql"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/elasticsearch"
	"github.com/qinsheng99/goWeb/library/mongo"
	"github.com/qinsheng99/goWeb/library/redisClient"
)

func Init(bundleDb *db.BundleDb, es *elasticsearch.ES, r *redisClient.Redis, mo *mongoClient.Mongo) (*Entry, error) {
	NewEsDao := persistence.NewEsDao(es)
	NewCustomerDao := customer.NewCustomerDao(bundleDb, NewEsDao)
	Drainage := drainage.NewDS(NewCustomerDao)
	NewHandlerDao := ceshi.NewHandler(NewCustomerDao, NewEsDao, Drainage)
	NewCeshi := ceshi2.NewCeshi()
	NewRedis := redisClient.NewRedisStruct(r)
	NewH := redis.NewH(NewCeshi, NewRedis)
	NewMysqlImp := ServiceMysql.NewMysqlService(bundleDb)
	NewPostgresBoy := boy.NewPostgresBoy()
	NewPostgresqlService := servicePostgresql.NewPostgresqlService(NewPostgresBoy)
	NewMysql := mysql.NewMysql(NewMysqlImp, NewPostgresqlService)
	NewSort := sortHandler.NewSort()
	NewEs := esHandle.NewEsHandle(NewEsDao, NewMysqlImp)
	NewMgoInterface := mongoClient.NewMongoStruct(mo)
	NewMgo := mongo.NewMgo(NewMgoInterface, NewRedis)
	NewDemo := demo.NewDemo(NewRedis)

	e := &Entry{
		NewHandler: NewHandlerDao,
		NewH:       NewH,
		NewMysql:   NewMysql,
		NewSort:    NewSort,
		NewEs:      NewEs,
		NewMgo:     NewMgo,
		NewDemo:    NewDemo,
	}

	return e, nil
}
