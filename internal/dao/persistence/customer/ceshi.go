package customer

import (
	"fmt"
	"time"

	drequest "github.com/qinsheng99/goWeb/api/entity/drainage/request"
	"github.com/qinsheng99/goWeb/internal/dao/idao"
	"github.com/qinsheng99/goWeb/internal/dao/idao/customer"
	"github.com/qinsheng99/goWeb/internal/model"
	"github.com/qinsheng99/goWeb/library/db"
	"github.com/qinsheng99/goWeb/library/funcTest"

	"gorm.io/gorm"
)

type CustomerDao struct {
	db *db.BundleDb
	es idao.EsImp
}

func NewCustomerDao(db *db.BundleDb, es idao.EsImp) customer.CustomerImp {
	return &CustomerDao{
		db: db,
		es: es,
	}
}

func (c *CustomerDao) GetList() (data []*model.CustomerFollowUser2, err error) {
	// err = c.Db.Db.Model(model.CustomerFollowUser1{}).
	// 	Table(model.CustomerFollowUser{}.TableName()).
	// 	Where("id > ?", 0).
	// 	Preload("Customer").
	// 	Preload("CustomerTags").
	// 	Preload("CustomerTransaction").
	// 	Limit(10).
	// 	Offset(0).
	// 	Find(&data).Error

	err = c.db.Db.Model(model.CustomerFollowUser{}).
		Where("id > ?", 0).
		Limit(10).
		Offset(0).
		Find(&data).Error
	return
}

func (c *CustomerDao) GetByIds(name string) (ids []int, err error) {

	err = c.db.Db.
		Model(&model.Customer{}).
		Where("name LIKE ?", "%"+name+"%").
		Scopes(customer.IsDelete()).
		Pluck("id", &ids).Error
	return
}

func (c *CustomerDao) GetData(companyId, gender int) (ids []int, err error) {

	err = c.db.Db.
		Model(&model.Customer{}).
		Where("gender = ?", gender).
		Where("companyId", companyId).
		Pluck("id", &ids).Error
	return
}

func (c *CustomerDao) Refresh() error {
	var customerFollowUserList []model.CustomerFollowUserList
	lastId := 0
	pageSize := 50

	begin := time.Now()
	// 分批插入es
	for {
		err := c.db.Db.Where("id > ?", lastId).Limit(pageSize).Order("id asc").Find(&customerFollowUserList).Error
		if err != nil {
			return err
		}

		var ids []int64
		for _, customerFollowUser := range customerFollowUserList {
			lastId = int(customerFollowUser.Id)
			ids = append(ids, customerFollowUser.Id)
		}

		// 查询并组装插入 es 的数据
		dbScope := func(db *gorm.DB) *gorm.DB {
			return db.Where("id in (?)", ids)
		}
		customerFollowerUserEs, err := c.getCustomerFollowerUserListForEs(dbScope)
		if err != nil {
			return err
		}

		// 批量插入 es
		c.es.InsertElastic(customerFollowerUserEs)

		if len(customerFollowUserList) < pageSize {
			break
		}
	}
	usedTime := time.Since(begin)
	fmt.Printf("es 数据刷新完成，耗时：%s", usedTime.String())

	return nil

}

func (c *CustomerDao) getCustomerFollowerUserListForEs(scope func(db *gorm.DB) *gorm.DB) ([]*model.CustomerFollowerUserEs, error) {
	var customerFollowerUsers []model.CustomerFollowUserWith
	query := c.db.Db.
		Scopes(scope).Model(model.CustomerFollowUser{}).
		Preload("Customer").
		Preload("CustomerTags").
		Preload("CustomerTransaction").
		Find(&customerFollowerUsers)

	if err := query.Error; err != nil {
		return nil, err
	}

	var returnData []*model.CustomerFollowerUserEs

	for _, customerFollowerUser := range customerFollowerUsers {
		returnData = append(returnData, &model.CustomerFollowerUserEs{
			Id:                  customerFollowerUser.Id,
			UserId:              customerFollowerUser.UserId,
			StaffId:             int(customerFollowerUser.StaffId),
			AddTime:             uint(customerFollowerUser.AddTime),
			AddChannel:          customerFollowerUser.AddChannal,
			DelFollow:           customerFollowerUser.DelFollow,
			DelExternal:         customerFollowerUser.DelExternal,
			Remark:              customerFollowerUser.Remark,
			CustomerId:          customerFollowerUser.CustomerId,
			Gender:              customerFollowerUser.Customer.Gender,
			CustomerName:        customerFollowerUser.Customer.Name,
			CustomerState:       funcTest.GetCustomerState(customerFollowerUser.DelFollow, customerFollowerUser.DelExternal),
			CustomerTags:        funcTest.GetTagIdListByTagList(customerFollowerUser.CustomerTags),
			LastTransactionTime: customerFollowerUser.CustomerTransaction.LastTransactionTime,
			VolumeTotal:         customerFollowerUser.CustomerTransaction.VolumeTotal,
			TransactionTotal:    customerFollowerUser.CustomerTransaction.TransactionTotal,
			CompanyId:           customerFollowerUser.CompanyId,
		})
	}

	return returnData, nil
}

func (c *CustomerDao) GetDrainageList(drainage drequest.DrainageRequest) (data []*model.QyDrainage, err error) {
	err = c.db.Db.Table(model.QyDrainage{}.TableName()).
		Preload("Staff").
		Preload("Code").
		Preload("Group").
		Where("type = ?", drainage.Type).
		Find(&data).Error
	return
}

func (c *CustomerDao) GetUsers(user []int) (data []*model.QyUser1, err error) {
	err = c.db.Db.Model(&model.QyUser{}).Where("staffId in ?", user).Find(&data).Error
	return
}

func (c *CustomerDao) GetTags(user []int) (data []*model.QyTag1, err error) {
	err = c.db.Db.Model(&model.QyTag{}).Where("id in ?", user).Find(&data).Error
	return
}
