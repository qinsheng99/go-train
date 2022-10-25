package ServiceMysql

import (
	"github.com/qinsheng99/go-train/internal/dao/idao/customer"
	"github.com/qinsheng99/go-train/internal/model"
	"github.com/qinsheng99/go-train/library/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlService struct {
	db *db.BundleDb
}

func NewMysqlService(db *db.BundleDb) customer.CeshiMysqlImp {
	return &MysqlService{
		db: db,
	}
}

func (m *MysqlService) GetCeshiData() (datas []*model.CeshiWith, err error) {
	if err = m.db.Db.Model(&model.CeshiWith{}).
		Preload("CeshiDemo", func(q *gorm.DB) *gorm.DB {
			return q.Where("id <> ?", 2)
		}).
		Where("back_name LIKE ?", "%修复%").
		Find(&datas).Error; err != nil {
		return nil, err
	}
	//for _, data := range datas {
	//	//data.ModifyTime = data.ModifyTime.Format("")
	//}
	return
}

func (m *MysqlService) GetFirstData() (data *model.CeshiDemo1, err error) {
	if err = m.db.Db.
		FirstOrInit(&data, model.CeshiDemo1{Id: 1}).Error; err != nil {
		return nil, err
	}
	return
}

func (m *MysqlService) InsertData(data *model.Ceshi, updateStringArr []string) (err error) {
	if err = m.db.Db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns(updateStringArr),
		}).Create(data).Error; err != nil {
		return err
	}
	return
}

func (m *MysqlService) JoinData() (datas []*model.CeshiJoin, err error) {
	err = m.db.Db.Model(&model.Ceshi{}).
		Select("ceshi.*, b.name as Name").
		Joins("left join ceshi_demo1 as b on b.uri = ceshi.uri").
		Joins("left join ceshi_demo2 as c on c.uri = ceshi.uri").
		Order(model.Ceshi{}.TableName() + ".id desc").
		Scan(&datas).Error
	return
}

func (m *MysqlService) UpdateData(id int, data string) (err error) {
	err = m.db.Db.Model(&model.Ceshi{}).Where("id = ?", id).Update("back_name", data).Error
	return
}

func (m *MysqlService) GetCeshiEsData() (datas []*model.CeshiEs, err error) {
	var pageSize, page, id, count = 10, 1, 0, 200
	var ceshi []*model.Ceshi
	for {
		err = m.db.Db.Model(&model.Ceshi{}).
			Where("id > ?", id).
			Limit(pageSize).
			Order("id asc").
			Find(&ceshi).Error
		if err != nil {
			return nil, err
		}
		for _, v := range ceshi {
			datas = append(datas, &model.CeshiEs{
				Id:         v.Id,
				Uri:        v.Uri,
				Tag:        v.Tag,
				BackName:   v.BackName,
				IsDelete:   v.IsDelete,
				CreateTime: v.CreateTime,
				DeleteTime: v.DeleteTime,
			})
			id = v.Id
		}
		page++

		if id == count {
			break
		}
	}
	return
}
