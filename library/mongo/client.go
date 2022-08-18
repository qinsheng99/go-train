package mongoClient

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mgo *mongo.Client

type Mongo struct {
	Mo *mongo.Client
}

func InitMongo() (*Mongo, error) {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	Mgo, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// 检查连接
	err = Mgo.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return &Mongo{Mo: Mgo}, nil
}
