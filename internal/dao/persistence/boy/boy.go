package boy

import (
	"fmt"

	"github.com/qinsheng99/go-train/internal/dao/idao/boy"
	"github.com/qinsheng99/go-train/internal/model"
	"github.com/qinsheng99/go-train/library/db"
	"gorm.io/gorm"
)

type boyDao struct {
	db *gorm.DB
}

func NewPostgresBoy() boy.Boyimpl {
	return &boyDao{db: db.GetPostgresqlDb()}
}

func (b *boyDao) Getlist() (data []*model.Boy, err error) {
	err = b.db.Model(&model.Boy{}).Find(&data).Error
	return
}

func (b *boyDao) GetAddress(col, s string) (data []*model.Boy, err error) {
	err = b.db.
		Model(&model.Boy{}).
		Select("*", fmt.Sprintf(`jsonb_path_query(information, '$.%s ? (@ starts with "%v")')`, col, s)).
		Find(&data).
		Error
	return
}

func (b *boyDao) GetOne(id int64) (data *model.Boy, err error) {
	err = b.db.
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
	err = b.db.
		Model(data).
		Select("name,id,information", fmt.Sprintf("arr[%d] as arrone", index)).
		Order("id desc").
		Find(data).
		Error
	return
}

func (b *boyDao) FindJson(q string, flag bool) (data []*model.Boy, err error) {
	query := b.db.Model(&model.Boy{}).Debug()
	if flag {
		query = query.Where("information @> ?::jsonb", "["+q+"]")
	} else {
		query = query.Where("information @> ?::jsonb", q)
	}

	err = query.Find(&data).Error
	return
}
