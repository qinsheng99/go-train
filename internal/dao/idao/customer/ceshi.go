package customer

import (
	drequest "gin/api/entity/drainage/request"
	"gin/internal/model"

	"gorm.io/gorm"
)

type CustomerImp interface {
	GetByIds(name string) (ids []int, err error)
	GetData(companyId, gender int) (ids []int, err error)
	Refresh() error
	GetList() (data []*model.CustomerFollowUser2, err error)
	GetDrainageList(drainage drequest.DrainageRequest) (data []*model.QyDrainage, err error)
	GetUsers(user []int) (data []*model.QyUser1, err error)
	GetTags(user []int) (data []*model.QyTag1, err error)
}

func IsDelete() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("isDelete", 0)
	}
}

type MysqlImp interface {
	GetCeshiData() ([]*model.CeshiWith, error)
	GetFirstData() (data *model.CeshiDemo1, err error)
	InsertData(data *model.Ceshi, updateStringArr []string) (err error)
	JoinData() (datas []*model.CeshiJoin, err error)
	UpdateData(id int, data string) (err error)
}
