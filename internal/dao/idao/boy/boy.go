package boy

import (
	"github.com/qinsheng99/goWeb/internal/model"
)

type Boyimpl interface {
	Getlist() ([]*model.Boy, error)
	GetAddress(string) ([]*model.Boy, error)
}

type BoyimplService interface {
	GetBoylist() ([]*model.Boy, error)
	GetBoyAddress(string) ([]*model.Boy, error)
}
