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

func (b *boyDao) GetOne(id int64) (data *model.Boy, err error) {
	err = db.GetPostgresqlDb().
		Model(&model.Boy{}).
		Where("id = ?", id).
		First(&data).
		Error
	return
}

func (b *boyDao) CreateOne(data *model.Boy) (_ *model.Boy, err error) {
	//err = db.GetPostgresqlDb().Model(data).Select("id").Last(data).Error
	//if err != nil {
	//	return nil, err
	//}
	//data.Id++
	err = db.GetPostgresqlDb().
		Model(&model.OldBoy{}).
		Create(&model.OldBoy{
			Name: data.Name,
			Arr:  data.Arr,
		}).Error
	return data, err
}
