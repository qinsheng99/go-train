package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id" json:"-"`
	Name string             `bson:"name" json:"name"`
	Age  int64              `bson:"age" json:"age"`
	Cve  string             `bson:"cve" json:"cve"`
	Dep  string             `bson:"dep" json:"dep"`
	Repo string             `bson:"repo" json:"repo"`
}
