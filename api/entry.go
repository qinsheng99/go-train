package api

import (
	"github.com/qinsheng99/goWeb/api/handel/ceshi"
	"github.com/qinsheng99/goWeb/api/handel/mysql"
	"github.com/qinsheng99/goWeb/api/handel/redis"
)

type Entry struct {
	NewHandler *ceshi.Handler
	NewH       *redis.Handle
	NewMysql   *mysql.Handler
	// NewCustomer customer.CustomerDao
}
