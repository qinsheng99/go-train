package api

import (
	"github.com/qinsheng99/goWeb/api/handel/ceshi"
	"github.com/qinsheng99/goWeb/api/handel/mysql"
	"github.com/qinsheng99/goWeb/api/handel/redis"
	sortHandler "github.com/qinsheng99/goWeb/api/handel/sort"
)

type Entry struct {
	NewHandler *ceshi.Handler
	NewH       *redis.Handle
	NewMysql   *mysql.Handler
	NewSort *sortHandler.SortHandler
	// NewCustomer customer.CustomerDao
}
