package mongo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	mongoRequest "github.com/qinsheng99/goWeb/api/entity/mongo"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/internal/model"
	mongoClient "github.com/qinsheng99/goWeb/library/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handle struct {
	mo mongoClient.MongoInterface
}

func NewMgo(mo mongoClient.MongoInterface) *Handle {
	return &Handle{mo: mo}
}

func (h *Handle) InsertOne(c *gin.Context) {
	var req mongoRequest.RequestOne
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.Failure(c, err)
		return
	}
	data := model.User{
		Id:   primitive.NewObjectID(),
		Name: req.DD.Name,
		Age:  req.DD.Age,
		Cve:  req.DD.Cve,
		Dep:  req.DD.Dep,
		Repo: req.DD.Repo,
	}

	in, err := h.mo.Collection("").InsertOne(context.Background(), data)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, in.InsertedID)
}

func (h *Handle) InsertMany(c *gin.Context) {
	var req mongoRequest.RequestMany
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.Failure(c, err)
		return
	}
	data := bson.A{}
	for _, v := range req.Data {
		data = append(data, model.User{
			Id:   primitive.NewObjectID(),
			Name: v.Name,
			Age:  v.Age,
			Cve:  v.Cve,
			Dep:  v.Dep,
			Repo: v.Repo,
		})
	}

	in, err := h.mo.Collection("").InsertMany(context.Background(), data)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, in.InsertedIDs)
}

func (h *Handle) Find(c *gin.Context) {
	name := c.QueryArray("name")

	filter := bson.M{
		"name": bson.M{
			"$in": name,
		},
	}

	var data []model.User
	err := h.mo.Collection("").Find(context.Background(), filter, &data)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, data)
}

func (h *Handle) FindOne(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		common.QueryFailure(c, nil)
		return
	}

	filter := bson.M{
		"name": name,
	}
	var data model.User
	err := h.mo.Collection("").FindOne(context.Background(), filter, &data, "age", "repo")
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, data)
}

func (h *Handle) Update(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		common.QueryFailure(c, nil)
		return
	}

	filter := bson.M{
		"name": name,
	}
	//update := bson.M{
	//	"$set": bson.M{
	//		"age": 24,
	//	},
	//}
	update := bson.M{
		"$inc": bson.M{
			"age": 1,
		},
	}
	_, err := h.mo.Collection("").Update(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, "")
}
