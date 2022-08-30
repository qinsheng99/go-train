package servicePostgresql

import (
	"encoding/json"

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

	for _, v := range data {
		var i model.Infor
		_ = json.Unmarshal(v.Informations, &i)
		v.Information = i
	}
	return
}

func (b *postgresqlService) GetBoyAddress(s string) (data []*model.Boy, err error) {
	data, err = b.boy.GetAddress(s)
	for _, v := range data {
		var i model.Infor
		_ = json.Unmarshal(v.Informations, &i)
		v.Information = i
	}
	return
}
