package service

import (
	"gin/internal/service/drainage"
	ServiceMysql "gin/internal/service/mysql"

	"github.com/google/wire"
)

var ServiceProvider = wire.NewSet(
	drainage.NewDS,
	ServiceMysql.NewMysqlService,
)
