package dao

import (
	"github.com/qinsheng99/go-train/internal/dao/persistence"
	"github.com/qinsheng99/go-train/internal/dao/persistence/boy"
	"github.com/qinsheng99/go-train/internal/dao/persistence/customer"

	"github.com/google/wire"
)

var DaoProvider = wire.NewSet(
	customer.NewCustomerDao,
	persistence.NewEsDao,
	boy.NewPostgresBoy,
)
