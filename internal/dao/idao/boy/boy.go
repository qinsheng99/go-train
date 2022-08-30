package boy

import (
	"github.com/qinsheng99/goWeb/internal/model"
)

type Boyimpl interface {
	Getlist() (data []model.Boy, err error)
}
