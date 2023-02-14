package mongoClient

import (
	"context"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStruct struct {
	mo         *Mongo
	collection *mongo.Collection
}

type Mi struct {
	I MongoInterface
	C MongoCondition
}

func NewMongoStruct(m *Mongo) Mi {
	mm := &MongoStruct{
		mo: m,
	}

	return Mi{
		I: mm,
		C: mm,
	}
}

func (m *MongoStruct) Collection(name string, opts ...*options.CollectionOptions) MongoInterface {
	if m.mo.database == nil {
		return m
	}
	if len(name) == 0 {
		m.collection = m.mo.database.Collection(m.mo.collection, opts...)
	} else {
		m.collection = m.mo.database.Collection(name, opts...)
	}

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

func (m *MongoStruct) InsertMany(ctx context.Context, document []interface{}, opts ...*options.InsertManyOptions) (insertOne *mongo.InsertManyResult, err error) {
	if m.collection == nil {
		err = errors.New("mongo collection is nil")
		return
	}
	insertOne, err = m.collection.InsertMany(ctx, document, opts...)

	return
}

func (m *MongoStruct) Find(ctx context.Context, filter interface{}, data interface{}, opts ...*options.FindOptions) (err error) {
	if m.collection == nil {
		err = errors.New("mongo collection is nil")
		return
	}
	var find *mongo.Cursor
	find, err = m.collection.Find(ctx, filter, opts...)

	if err != nil {
		return
	}

	if err = m.validation(data, true); err != nil {
		return
	}

	if err = m.scanf(ctx, data, find); err != nil {
		return
	}

	return
}

func (m *MongoStruct) scanf(ctx context.Context, data interface{}, find *mongo.Cursor) error {
	return find.All(ctx, data)
}

func (m *MongoStruct) validation(data interface{}, flag bool) error {
	tyof := reflect.TypeOf(data)
	if tyof.Kind() != reflect.Ptr {
		return errors.New("data is not ptr")
	}
	if flag {
		tyof = tyof.Elem()
		if tyof.Kind() != reflect.Slice {
			return errors.New("data is not slice")
		}
	}

	return nil
}

func (m *MongoStruct) FindOne(ctx context.Context, filter interface{}, data interface{}, project bson.M) (err error) {
	if m.collection == nil {
		err = errors.New("mongo collection is nil")
		return
	}
	var find *mongo.SingleResult

	find = m.collection.FindOne(ctx, filter, &options.FindOneOptions{Projection: project})

	if err != nil {
		return
	}

	if err = m.validation(data, false); err != nil {
		return
	}

	if err = find.Decode(data); err != nil {
		return
	}

	return
}

func (m *MongoStruct) findOneOpt(column ...string) bson.M {
	if len(column) == 0 {
		return nil
	}
	var result = make(bson.M)
	for _, v := range column {
		result[v] = 1
	}

	return result
}

func (m *MongoStruct) Update(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (updates *mongo.UpdateResult, err error) {
	if m.collection == nil {
		err = errors.New("mongo collection is nil")
		return
	}
	updates, err = m.collection.UpdateMany(ctx, filter, update, opts...)

	return
}

func (m *MongoStruct) UpdatePush(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (updates *mongo.UpdateResult, err error) {
	if m.collection == nil {
		err = errors.New("mongo collection is nil")
		return
	}
	updates, err = m.collection.UpdateMany(ctx, filter, update, opts...)

	return
}

func (m *MongoStruct) Aggregate(ctx context.Context, pipeline interface{}, data interface{}, opts ...*options.AggregateOptions) (err error) {
	if m.collection == nil {
		err = errors.New("mongo collection is nil")
		return
	}
	var find *mongo.Cursor
	find, err = m.collection.Aggregate(ctx, pipeline, opts...)

	if err != nil {
		return
	}

	if err = m.validation(data, true); err != nil {
		return
	}

	if err = m.scanf(ctx, data, find); err != nil {
		return
	}

	return
}
