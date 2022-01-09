package dao

import (
	"github.com/qinsheng99/goWeb/internal/dao/persistence/customer"

	"github.com/google/wire"
)

var DaoProvider = wire.NewSet(
	customer.NewCustomerDao,
)
