package mongoClient

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongos interface {
	MongoInterface
	MongoCondition
}

type MongoInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) Mongos

	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (insertOne *mongo.InsertOneResult, err error)
	InsertMany(ctx context.Context, document []interface{}, opts ...*options.InsertManyOptions) (insertOne *mongo.InsertManyResult, err error)

	Find(ctx context.Context, filter interface{}, data interface{}, opts ...*options.FindOptions) (err error)
	FindOne(ctx context.Context, filter interface{}, data interface{}, project bson.M) (err error)

	Update(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (updates *mongo.UpdateResult, err error)
}

type MongoCondition interface {
	FieldIn(field bson.M, column string, data []string) bson.M
	FilterOrChooseColumn(field bson.M, flag bool, column ...string) bson.M
	FieldInc(field bson.M, column string, data interface{}) bson.M
	FieldSet(field bson.M, column string, data interface{}) bson.M
	SetOrInsert(data bson.M) bson.M
	Filter(data []Filter) bson.M
}
