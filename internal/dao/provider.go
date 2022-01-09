package dao

import (
	"gin/internal/dao/persistence/customer"

	"github.com/google/wire"
)

var DaoProvider = wire.NewSet(
	customer.NewCustomerDao,
)
