//+build wireinject

package api

import (
	"gin/api/handel"
	"gin/internal"
	"gin/library/db"
	"gin/library/elasticsearch"
	"gin/library/redisClient"

	"github.com/google/wire"
)

func Init(
	bundleDb *db.BundleDb, es *elasticsearch.ES,r *redisClient.Redis,
) (*Entry, error) {
	panic(wire.Build(
		internal.InternalProvider,
		handel.HandlerProvider,
		wire.Struct(new(Entry), "*"),
	))
}
