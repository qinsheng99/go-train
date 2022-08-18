package mongo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	mongoRequest "github.com/qinsheng99/goWeb/api/entity/mongo"
	"github.com/qinsheng99/goWeb/api/tools/common"
	"github.com/qinsheng99/goWeb/internal/model"
	mongoClient "github.com/qinsheng99/goWeb/library/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handle struct {
	mo mongoClient.MongoInterface
}

func NewMgo(mo mongoClient.MongoInterface) *Handle {
	return &Handle{mo: mo}
}

func (h *Handle) InsertOne(c *gin.Context) {
	var req mongoRequest.Request
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.Failure(c, err)
		return
	}
	data := model.User{
		Id:   primitive.NewObjectID(),
		Name: req.Name,
		Age:  req.Age,
		Cve:  req.Cve,
		Dep:  req.Dep,
		Repo: req.Repo,
	}

	in, err := h.mo.InsertOne(context.Background(), data)
	if err != nil {
		common.Failure(c, err)
		return
	}

	common.Success(c, in.InsertedID)
}
