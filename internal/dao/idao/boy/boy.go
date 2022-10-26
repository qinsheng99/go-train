package boy

import (
	postgresqlRequest "github.com/qinsheng99/go-train/api/entity/postgresql"
	"github.com/qinsheng99/go-train/internal/model"
)

type Boyimpl interface {
	Getlist() ([]*model.Boy, error)
	GetAddress(string, string) ([]*model.Boy, error)

	GetOne(int64) (*model.Boy, error)

	CreateOne(*model.Boy) (*model.Boy, error)

	FindArrOne(int64, interface{}) error
	FindJson(string, bool) ([]*model.Boy, error)
}

type BoyimplService interface {
	GetBoylist() ([]*model.Boy, error)
	GetBoyAddress(string, string) ([]*model.Boy, error)
	GetBoyOne(int64) (*model.Boy, error)
	FindArrOne(int64) ([]model.BoyArr, error)

	CreateOne(postgresqlRequest.Boy) (*model.Boy, error)

	FindJson(string, bool) ([]*model.Boy, error)
}
