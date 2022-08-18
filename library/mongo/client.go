package mongoClient

import (
	"context"
	"fmt"

	"github.com/qinsheng99/goWeb/library/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mgo *mongo.Client

type Mongo struct {
	mo         *mongo.Client
	database   *mongo.Database
	collection string
}

func InitMongo(cfg *config.MongoConfig) (*Mongo, error) {
	var err error
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port))

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

	return &Mongo{mo: Mgo, database: Mgo.Database(cfg.Database), collection: cfg.Collection}, nil
}
