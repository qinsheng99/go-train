package boy

import (
	postgresqlRequest "github.com/qinsheng99/go-train/api/entity/postgresql"
	"github.com/qinsheng99/go-train/internal/model"
)

type Boyimpl interface {
	Getlist() ([]*model.Boy, error)
	GetAddress(string, string) ([]*model.Boy, error)

	GetOne(id int64) (data *model.Boy, err error)

	CreateOne(data *model.Boy) (_ *model.Boy, err error)

	FindArrOne(index int64, data interface{}) (err error)
}

type BoyimplService interface {
	GetBoylist() ([]*model.Boy, error)
	GetBoyAddress(string, string) ([]*model.Boy, error)
	GetBoyOne(id int64) (data *model.Boy, err error)
	FindArrOne(index int64) (data []model.BoyArr, err error)

	CreateOne(data postgresqlRequest.Boy) (_ *model.Boy, err error)
}
