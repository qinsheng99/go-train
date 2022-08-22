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

// FieldInc $inc可以对文档的某个值为数字型（只能为满足要求的数字）的键进行增减的操作。
func (m *MongoStruct) FieldInc(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$inc"] = bson.M{
		column: data,
	}
	return field
}

// FieldSet 用来指定一个键并更新键值，若键不存在并创建 { $set : { field : value } }
func (m *MongoStruct) FieldSet(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$set"] = bson.M{
		column: data,
	}
	return field
}

// FieldUnSet 用来删除一个键  { $unset : { field : 1} }
func (m *MongoStruct) FieldUnSet(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$unset"] = bson.M{
		column: data,
	}
	return field
}

// FieldPush 把value追加到field里面去，field一定要是数组类型才行，如果field不存在，会新增一个数组类型加进去 { $push : { field : value } }
func (m *MongoStruct) FieldPush(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$push"] = bson.M{
		column: data,
	}
	return field
}

// FieldPushAll 同$push,只是一次可以追加多个值到一个数组字段内 { $pushAll : { field : value_array } }
func (m *MongoStruct) FieldPushAll(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$pushAll"] = bson.M{
		column: data,
	}
	return field
}

// FieldPull 从数组field内删除一个等于value值  { $pull : { field : _value } }
func (m *MongoStruct) FieldPull(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$pull"] = bson.M{
		column: data,
	}
	return field
}

// FieldRename 修改字段名称 { $rename : { old_field_name : new_field_name } }
func (m *MongoStruct) FieldRename(field bson.M, column string, data interface{}) bson.M {
	if field == nil {
		field = make(bson.M)
	}
	field["$rename"] = bson.M{
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
