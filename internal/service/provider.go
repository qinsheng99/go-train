package service

import (
	"github.com/qinsheng99/goWeb/internal/service/drainage"
	ServiceMysql "github.com/qinsheng99/goWeb/internal/service/mysql"
	servicePostgresql "github.com/qinsheng99/goWeb/internal/service/postgresql"

	"github.com/google/wire"
)

var ServiceProvider = wire.NewSet(
	drainage.NewDS,
	ServiceMysql.NewMysqlService,
	servicePostgresql.NewPostgresqlService,
)
