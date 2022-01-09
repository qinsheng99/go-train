package internal

import (
	"gin/internal/dao"
	"gin/internal/service"

	"github.com/google/wire"
)

var InternalProvider = wire.NewSet(
	dao.DaoProvider,
	service.ServiceProvider,
)
