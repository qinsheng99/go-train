package handel

import (
	"github.com/qinsheng99/goWeb/api/handel/ceshi"
	"github.com/qinsheng99/goWeb/api/handel/redis"

	"github.com/google/wire"
)

var HandlerProvider = wire.NewSet(
	ceshi.NewHandler,
	redis.NewH,
)
