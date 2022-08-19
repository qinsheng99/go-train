package mongoClient

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Filter struct {
	Column string
	Data   interface{}
}

func (m *MongoStruct) FieldIn(field bson.M, column string, data []string) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field[column] = bson.M{"$in": data}
	return field
}

func (m *MongoStruct) FilterOrChooseColumn(field bson.M, flag bool, column ...string) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	if flag {
		for _, v := range column {
			field[v] = 1
		}
	} else {
		for _, v := range column {
			field[v] = 0
		}
	}
	return field
}

func (m *MongoStruct) FieldInc(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$inc"] = bson.M{
		column: data,
	}
	return field
}

func (m *MongoStruct) FieldSet(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$set"] = bson.M{
		column: data,
	}
	return field
}

func (m *MongoStruct) SetOrInsert(data bson.M) bson.M {
	return bson.M{
		"$setOnInsert": data,
	}
}

func (m *MongoStruct) Filter(data []Filter) bson.M {
	var result = make(bson.M)
	for _, v := range data {
		result[v.Column] = v.Data
	}
	return result
}
