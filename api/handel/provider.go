package handel

import (
	"github.com/qinsheng99/go-train/api/handel/ceshi"
	"github.com/qinsheng99/go-train/api/handel/demo"
	esHandle "github.com/qinsheng99/go-train/api/handel/es"
	"github.com/qinsheng99/go-train/api/handel/mongo"
	"github.com/qinsheng99/go-train/api/handel/mysql"
	"github.com/qinsheng99/go-train/api/handel/postgresql"
	"github.com/qinsheng99/go-train/api/handel/redis"
	sortHandler "github.com/qinsheng99/go-train/api/handel/sort"

	"github.com/google/wire"
)

var HandlerProvider = wire.NewSet(
	ceshi.NewHandler,
	redis.NewH,
	esHandle.NewEsHandle,
	mysql.NewMysql,
	sortHandler.NewSort,
	mongo.NewMgo,
	postgresql.NewPostgreSql,
	demo.NewDemo,
)
