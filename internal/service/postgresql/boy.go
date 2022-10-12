package servicePostgresql

import (
	"github.com/jackc/pgtype"
	postgresqlRequest "github.com/qinsheng99/goWeb/api/entity/postgresql"
	"github.com/qinsheng99/goWeb/internal/dao/idao/boy"
	"github.com/qinsheng99/goWeb/internal/model"
)

type postgresqlService struct {
	boy boy.Boyimpl
}

func NewPostgresqlService(boy boy.Boyimpl) boy.BoyimplService {
	return &postgresqlService{boy: boy}
}

func (b *postgresqlService) GetBoylist() (data []*model.Boy, err error) {
	data, err = b.boy.Getlist()
	if err != nil {
		return nil, err
	}

	//for _, v := range data {
	//	var i model.Infor
	//	_ = json.Unmarshal(v.Informations, &i)
	//	v.Information = i
	//}
	return
}

func (b *postgresqlService) GetBoyAddress(s string) (data []*model.Boy, err error) {
	data, err = b.boy.GetAddress(s)
	//for _, v := range data {
	//	var i model.Infor
	//	_ = json.Unmarshal(v.Informations, &i)
	//	v.Information = i
	//}
	return
}

func (b *postgresqlService) GetBoyOne(id int64) (data *model.Boy, err error) {
	data, err = b.boy.GetOne(id)
	//var i model.Infor
	//_ = json.Unmarshal(data.Informations, &i)
	//data.Information = i
	return
}

func (b *postgresqlService) CreateOne(req postgresqlRequest.Boy) (data *model.Boy, err error) {
	info := pgtype.JSONB{}
	if err = info.Set(req.Informations); err != nil {
		return nil, err
	}
	data, err = b.boy.CreateOne(&model.Boy{
		Name:         req.Name,
		Informations: info,
		Arr:          req.Arr,
	})
	//var i model.Infor
	//_ = json.Unmarshal(data.Informations, &i)
	//data.Information = i
	return data, err
}
