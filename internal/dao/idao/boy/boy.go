package boy

import (
	postgresqlRequest "github.com/qinsheng99/goWeb/api/entity/postgresql"
	"github.com/qinsheng99/goWeb/internal/model"
)

type Boyimpl interface {
	Getlist() ([]*model.Boy, error)
	GetAddress(string) ([]*model.Boy, error)

	GetOne(id int64) (data *model.Boy, err error)

	CreateOne(data *model.Boy) (_ *model.Boy, err error)
}

type BoyimplService interface {
	GetBoylist() ([]*model.Boy, error)
	GetBoyAddress(string) ([]*model.Boy, error)

	GetBoyOne(id int64) (data *model.Boy, err error)

	CreateOne(data postgresqlRequest.Boy) (_ *model.Boy, err error)
}
