package boy

import (
	"fmt"

	"github.com/qinsheng99/go-train/internal/dao/idao/boy"
	"github.com/qinsheng99/go-train/internal/model"
	"github.com/qinsheng99/go-train/library/db"
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

func (b *boyDao) GetAddress(col, s string) (data []*model.Boy, err error) {
	err = db.GetPostgresqlDb().
		Model(&model.Boy{}).
		Select("*", fmt.Sprintf(`jsonb_path_query(information, '$.%s ? (@ starts with "%v")')`, col, s)).
		Find(&data).
		Error
	return
}

func (b *boyDao) GetOne(id int64) (data *model.Boy, err error) {
	err = db.GetPostgresqlDb().
		Model(&model.Boy{}).
		Where("id = ?", id).
		First(&data).
		Error
	return
}

func (b *boyDao) CreateOne(data *model.Boy) (_ *model.Boy, err error) {
	return data.Insert()
}

func (b *boyDao) FindArrOne(index int64, data interface{}) (err error) {
	err = db.GetPostgresqlDb().
		Model(data).
		Select("name,id,information", fmt.Sprintf("arr[%d] as arrone", index)).
		Order("id desc").
		Find(data).
		Error
	return
}
