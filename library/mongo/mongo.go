package mongoClient

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStruct struct {
	mo         *Mongo
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoStruct(r *Mongo) MongoInterface {
	return &MongoStruct{
		mo: r,
	}
}

func (m *MongoStruct) Database(name string, opts ...*options.DatabaseOptions) MongoInterface {
	m.database = m.mo.Mo.Database(name, opts...)
	return m
}

func (m *MongoStruct) Collection(name string, opts ...*options.CollectionOptions) MongoInterface {
	if m.database == nil {
		return m
	}
	m.collection = m.database.Collection(name, opts...)
	return m
}

func (m *MongoStruct) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (insertOne *mongo.InsertOneResult, err error) {
	if m.collection == nil {
		err = errors.New("mongo collection is nil")
		return
	}
	insertOne, err = m.collection.InsertOne(ctx, document, opts...)
	return
}
