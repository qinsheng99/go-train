package drainage

import (
	"strings"

	drequest "github.com/qinsheng99/go-train/api/entity/drainage/request"
	"github.com/qinsheng99/go-train/api/entity/drainage/response"
	"github.com/qinsheng99/go-train/internal/dao/idao/customer"
	"github.com/qinsheng99/go-train/internal/model"
	"github.com/qinsheng99/go-train/library/funcTest"
)

type Drainage interface {
	GetListForDrainage(d drequest.DrainageRequest) (data []*response.DrainageResponse, err error)
}

type DS struct {
	customer customer.CustomerImp
}

func NewDS(customer customer.CustomerImp) Drainage {
	return &DS{
		customer: customer,
	}
}

func (ds *DS) GetListForDrainage(d drequest.DrainageRequest) (data []*response.DrainageResponse, err error) {
	res, err := ds.customer.GetDrainageList(d)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		tags, e := ds.customer.GetTags(funcTest.StringToIint(strings.Split(v.TagIds, ",")))
		if e != nil {
			tags = []*model.QyTag1{}
		}
		users, e := ds.customer.GetUsers(funcTest.StringToIint(strings.Split(v.UserIds, ",")))
		if e != nil {
			users = []*model.QyUser1{}
		}
		data = append(data, &response.DrainageResponse{
			Name:      v.LinkName,
			BackType:  v.BackType,
			BackImage: v.BackImage,
			BrandName: v.BrandName,
			BrandLogo: v.BrandLogo,
			Users:     users,
			Tags:      tags,
			QrCode:    v.Code.QrCode,
		})
	}
	return data, nil
}
