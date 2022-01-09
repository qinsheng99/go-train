package internal

import (
	"github.com/qinsheng99/goWeb/internal/dao"
	"github.com/qinsheng99/goWeb/internal/service"

	"github.com/google/wire"
)

var InternalProvider = wire.NewSet(
	dao.DaoProvider,
	service.ServiceProvider,
)
