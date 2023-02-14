package mongo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongoRequest "github.com/qinsheng99/go-train/api/entity/mongo"
	"github.com/qinsheng99/go-train/api/tools/common"
	"github.com/qinsheng99/go-train/internal/model"
	mongoClient "github.com/qinsheng99/go-train/library/mongo"
	"github.com/qinsheng99/go-train/library/redisClient"
)

type Handle struct {
	mo    mongoClient.Mi
	redis redisClient.RedisInterface
}

const key = "insert_%v"

func NewMgo(mo mongoClient.Mi, redis redisClient.RedisInterface) *Handle {
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

	in, err := h.mo.I.Collection("cla").InsertOne(context.Background(), data)
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

	in, err := h.mo.I.Collection("").InsertMany(context.Background(), data)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, in.InsertedIDs)
}

func (h *Handle) Find(c *gin.Context) {
	name := c.QueryArray("name")

	var data []model.User
	err := h.mo.I.Collection("").Find(
		context.Background(),
		h.mo.C.FieldIn(nil, "name", name),
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

	var data model.DCLA
	err := h.mo.I.Collection("cla").FindOne(
		context.Background(),
		h.mo.C.Filter([]mongoClient.Filter{{Column: "url", Data: name}}),
		&data,
		nil)
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
	_, err := h.mo.I.Collection("").Update(
		context.Background(),
		h.mo.C.Filter([]mongoClient.Filter{{Column: "name", Data: name}}),
		//h.mo.FieldInc(nil, "age", 1),
		h.mo.C.FieldSet(nil, "age", 27),
	)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, "")
}

func (h *Handle) Push(c *gin.Context) {
	name, ok := c.GetQuery("url")
	if !ok {
		common.QueryFailure(c, nil)
		return
	}
	_, err := h.mo.I.Collection("cla").Update(
		context.Background(),
		h.mo.C.Filter([]mongoClient.Filter{{Column: "url", Data: name}}),
		//h.mo.FieldInc(nil, "age", 1),
		h.mo.C.FieldPush(nil, "fields", []model.DField{
			{
				ID:          "ID1",
				Title:       "Title1",
				Type:        "Type1",
				Description: "Description1",
				Required:    false,
			},
		}),
	)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, "")
}

func (h *Handle) InsertWukong(c *gin.Context) {
	var d model.DWuKong
	k := c.Query("key")

	d.Id = k + "qwe"

	d.Samples = []model.DSample{
		{Num: 1, Name: k + "1"},
		{Num: 2, Name: k + "2"},
		{Num: 3, Name: k + "3"},
		{Num: 4, Name: k + "4"},
		{Num: 5, Name: k + "5"},
	}

	d.Pictures = []model.DictureInfo{
		{Desc: k + "1", Link: "1", Style: "1"},
		{Desc: k + "2", Link: "2", Style: "2"},
		{Desc: k + "3", Link: "3", Style: "3"},
		{Desc: k + "4", Link: "4", Style: "4"},
		{Desc: k + "5", Link: "5", Style: "5"},
		{Desc: k + "6", Link: "6", Style: "6"},
	}

	one, err := h.mo.I.Collection("wukong").InsertOne(context.Background(), d)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, one.InsertedID)
}

func (h *Handle) FindWukong(c *gin.Context) {
	//id := c.Query("id")

	var res []model.DWuKong

	var f = new(options.FindOptions)
	f.SetLimit(2)
	f.SetSkip(0)

	err := h.mo.I.Collection("wukong").
		Find(context.Background(), bson.M{}, &res, f)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, res)
}

func (h *Handle) AggregateWukong(c *gin.Context) {
	var req = struct {
		Page int
		Size int
		Id   string
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		common.QueryFailure(c, err)
		return
	}

	fieldRef := "$pictures"

	project := bson.M{
		"pictures": bson.M{"$slice": bson.A{
			fieldRef, (req.Page - 1) * req.Size, req.Size,
		}},
		"total": bson.M{
			"$cond": bson.M{
				"if":   bson.M{"$isArray": fieldRef},
				"then": bson.M{"$size": fieldRef},
				"else": 0,
			},
		},
	}

	pipeline := bson.A{
		bson.M{"$match": bson.M{"id": req.Id}},
		bson.M{"$project": project},
	}

	var v []struct {
		Total    int                 `bson:"total"`
		Pictures []model.DictureInfo `bson:"pictures"`
	}

	err := h.mo.I.Collection("wukong").Aggregate(context.Background(), pipeline, &v)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, v)
}
