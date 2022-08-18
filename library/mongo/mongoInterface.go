package mongoClient

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInterface interface {
	Database(name string, opts ...*options.DatabaseOptions) MongoInterface
	Collection(name string, opts ...*options.CollectionOptions) MongoInterface

	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (insertOne *mongo.InsertOneResult, err error)
}
