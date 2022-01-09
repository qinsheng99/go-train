package api

import (
	"gin/api/handel/ceshi"
	"gin/api/handel/mysql"
	"gin/api/handel/redis"
)

type Entry struct {
	NewHandler *ceshi.Handler
	NewH       *redis.Handle
	NewMysql   *mysql.Handler
	// NewCustomer customer.CustomerDao
}
