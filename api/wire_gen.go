//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"github.com/qinsheng99/go-train/api/handel/ceshi"
	"github.com/qinsheng99/go-train/api/handel/demo"
	esHandle "github.com/qinsheng99/go-train/api/handel/es"
	"github.com/qinsheng99/go-train/api/handel/mongo"
	"github.com/qinsheng99/go-train/api/handel/mysql"
	"github.com/qinsheng99/go-train/api/handel/postgresql"
	"github.com/qinsheng99/go-train/api/handel/redis"
	sortHandler "github.com/qinsheng99/go-train/api/handel/sort"
	"github.com/qinsheng99/go-train/internal/dao/persistence"
	"github.com/qinsheng99/go-train/internal/dao/persistence/boy"
	"github.com/qinsheng99/go-train/internal/dao/persistence/customer"
	ceshi2 "github.com/qinsheng99/go-train/internal/service/ceshi"
	"github.com/qinsheng99/go-train/internal/service/drainage"
	ServiceMysql "github.com/qinsheng99/go-train/internal/service/mysql"
	servicePostgresql "github.com/qinsheng99/go-train/internal/service/postgresql"
	"github.com/qinsheng99/go-train/library/db"
	"github.com/qinsheng99/go-train/library/elasticsearch"
	mongoClient "github.com/qinsheng99/go-train/library/mongo"

	"github.com/qinsheng99/go-train/library/redisClient"
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
	NewMysql := mysql.NewMysql(NewMysqlImp)
	NewPostgreSql := postgresql.NewPostgreSql(NewPostgresqlService)
	NewSort := sortHandler.NewSort()
	NewEs := esHandle.NewEsHandle(NewEsDao, NewMysqlImp)
	NewMgoInterface := mongoClient.NewMongoStruct(mo)
	NewMgo := mongo.NewMgo(NewMgoInterface, NewRedis)
	NewDemo := demo.NewDemo(NewRedis)

	e := &Entry{
		NewHandler:    NewHandlerDao,
		NewH:          NewH,
		NewMysql:      NewMysql,
		NewSort:       NewSort,
		NewEs:         NewEs,
		NewMgo:        NewMgo,
		NewDemo:       NewDemo,
		NewPostgreSql: NewPostgreSql,
	}

	return e, nil
}
