package mongo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	mongoRequest "github.com/qinsheng99/goWeb/api/entity/mongo"
	"github.com/qinsheng99/goWeb/api/tools"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/internal/model"
	mongoClient "github.com/qinsheng99/goWeb/library/mongo"
	"github.com/qinsheng99/goWeb/library/redisClient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handle struct {
	mo    mongoClient.Mongos
	redis redisClient.RedisInterface
}

const key = "insert_%v"

func NewMgo(mo mongoClient.Mongos, redis redisClient.RedisInterface) *Handle {
	return &Handle{mo: mo, redis: redis}
}

// InsertOne @Summary Create
// @Description create project
// @Tags  Project
// @Param	RequestOne 	true	"body of creating project"
// @Accept json
// @Produce json
// @Router /mongo/insert-one [post]
func (h *Handle) InsertOne(c *gin.Context) {
	var req mongoRequest.RequestOne
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.Failure(c, err)
		return
	}
	data := model.DCLA{
		URL:          "123",
		Text:         "text",
		OrgSignature: []byte("ceshi"),
		DCLAInfo: model.DCLAInfo{
			Fields: []model.DField{
				{
					ID:          "ID",
					Title:       "Title",
					Type:        "Type",
					Description: "Description",
					Required:    false,
				},
			},
			Language:         "Language",
			CLAHash:          "CLAHash",
			OrgSignatureHash: "OrgSignatureHash",
		},
	}

	//in, err := h.mo.Collection("cla").InsertOne(context.Background(), data)
	//if err != nil {
	//	common.Failure(c, err)
	//	return
	//}
	//
	//common.Success(c, in.InsertedID)
	toMap, _ := tools.StructToMap(data)
	common.Success(c, toMap)
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

	var data []model.User
	err := h.mo.Collection("").Find(
		context.Background(),
		h.mo.FieldIn(nil, "name", name),
		&data,
	)
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

	var data model.User
	err := h.mo.Collection("").FindOne(
		context.Background(),
		h.mo.Filter([]mongoClient.Filter{{Column: "name", Data: name}}),
		&data,
		h.mo.FilterOrChooseColumn(nil, true, "age", "repo"))
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
	_, err := h.mo.Collection("").Update(
		context.Background(),
		h.mo.Filter([]mongoClient.Filter{{Column: "name", Data: name}}),
		h.mo.FieldInc(nil, "age", 1),
	)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, "")
}
