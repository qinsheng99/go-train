package mongoClient

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) MongoInterface

	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (insertOne *mongo.InsertOneResult, err error)
	InsertMany(ctx context.Context, document []interface{}, opts ...*options.InsertManyOptions) (insertOne *mongo.InsertManyResult, err error)

	Find(ctx context.Context, filter interface{}, data interface{}, opts ...*options.FindOptions) (err error)
	FindOne(ctx context.Context, filter interface{}, data interface{}, column ...string) (err error)

	Update(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (updates *mongo.UpdateResult, err error)
}
