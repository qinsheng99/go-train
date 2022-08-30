package boy

import (
	"fmt"

	"github.com/qinsheng99/goWeb/internal/dao/idao/boy"
	"github.com/qinsheng99/goWeb/internal/model"
	"github.com/qinsheng99/goWeb/library/db"
)

type boyDao struct {
}

func NewPostgresBoy() boy.Boyimpl {
	return &boyDao{}
}

func (b *boyDao) Getlist() (data []*model.Boy, err error) {
	err = db.GetPostgresqlDb().Model(&model.Boy{}).Find(&data).Error
	return
}

func (b *boyDao) GetAddress(s string) (data []*model.Boy, err error) {
	err = db.GetPostgresqlDb().
		Model(&model.Boy{}).
		Select("*", fmt.Sprintf(`jsonb_path_query(information, '$.address ? (@ starts with "%v")')`, s)).
		Find(&data).
		Error
	return
}
