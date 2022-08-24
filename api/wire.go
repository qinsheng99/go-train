//go:build wireinject
// +build wireinject

package api

import (
	"github.com/qinsheng99/goWeb/api/handel"
	"github.com/qinsheng99/goWeb/internal"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/elasticsearch"
	"github.com/qinsheng99/goWeb/library/mongo"
	"github.com/qinsheng99/goWeb/library/redisClient"

	"github.com/google/wire"
)

func Init(
	bundleDb *db.BundleDb, es *elasticsearch.ES, r *redisClient.Redis, mo *mongoClient.Mongo,
) (*Entry, error) {
	panic(wire.Build(
		internal.InternalProvider,
		handel.HandlerProvider,
		wire.Struct(new(Entry), "*"),
	))
}
