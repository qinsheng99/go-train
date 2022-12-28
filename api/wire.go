//go:build wireinject
// +build wireinject

package api

import (
	"github.com/qinsheng99/go-train/api/handel"
	"github.com/qinsheng99/go-train/internal"
	"github.com/qinsheng99/go-train/library"
	"github.com/qinsheng99/go-train/library/db"
	"github.com/qinsheng99/go-train/library/elasticsearch"
	"github.com/qinsheng99/go-train/library/mongo"
	"github.com/qinsheng99/go-train/library/redisClient"

	"github.com/google/wire"
)

func Init(
	bundleDb *db.BundleDb, es *elasticsearch.ES, r *redisClient.Redis, mo *mongoClient.Mongo, post *db.BundlePostgresql,
) (*Entry, error) {
	panic(wire.Build(
		internal.InternalProvider,
		handel.HandlerProvider,
		library.LibraryProvider,
		wire.Struct(new(Entry), "*"),
	))
}
