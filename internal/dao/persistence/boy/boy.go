package boy

import (
	"github.com/qinsheng99/goWeb/internal/dao/idao/boy"
	"github.com/qinsheng99/goWeb/internal/model"
	"github.com/qinsheng99/goWeb/library/db"
)

type boyDao struct {
}

func NewPostgresBoy() boy.Boyimpl {
	return &boyDao{}
}

func (b *boyDao) Getlist() (data []model.Boy, err error) {
	err = db.GetPostgresqlDb().Model(&model.Boy{}).Find(&data).Error
	return
}
