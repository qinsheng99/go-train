package api

import (
	"github.com/qinsheng99/goWeb/api/handel/ceshi"
	"github.com/qinsheng99/goWeb/api/handel/demo"
	esHandle "github.com/qinsheng99/goWeb/api/handel/es"
	"github.com/qinsheng99/goWeb/api/handel/mongo"
	"github.com/qinsheng99/goWeb/api/handel/mysql"
	"github.com/qinsheng99/goWeb/api/handel/postgresql"
	"github.com/qinsheng99/goWeb/api/handel/redis"
	sortHandler "github.com/qinsheng99/goWeb/api/handel/sort"
)

type Entry struct {
	NewHandler    *ceshi.Handler
	NewH          *redis.Handle
	NewMysql      *mysql.Handler
	NewPostgreSql *postgresql.Handler
	NewSort       *sortHandler.SortHandler
	NewEs         *esHandle.EsHandle
	NewMgo        *mongo.Handle
	NewDemo       *demo.Handle
	// NewCustomer customer.CustomerDao
}
