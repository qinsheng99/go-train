package internal

import (
	"github.com/qinsheng99/go-train/internal/dao"
	"github.com/qinsheng99/go-train/internal/service"

	"github.com/google/wire"
)

var InternalProvider = wire.NewSet(
	dao.DaoProvider,
	service.ServiceProvider,
)
