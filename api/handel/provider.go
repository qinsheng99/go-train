package handel

import (
	"gin/api/handel/ceshi"
	"gin/api/handel/redis"

	"github.com/google/wire"
)

var HandlerProvider = wire.NewSet(
	ceshi.NewHandler,
	redis.NewH,
)
