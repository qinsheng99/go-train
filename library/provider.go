package library

import (
	"github.com/google/wire"

	mongoClient "github.com/qinsheng99/go-train/library/mongo"
	"github.com/qinsheng99/go-train/library/redisClient"
)

var LibraryProvider = wire.NewSet(
	mongoClient.NewMongoStruct,
	redisClient.NewRedisStruct,
)
