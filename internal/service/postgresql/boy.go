package servicePostgresql

import (
	"encoding/json"

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
	return
}

func (b *postgresqlService) GetBoyAddress(s string) (data []*model.Boy, err error) {
	data, err = b.boy.GetAddress(s)
	return
}

func (b *postgresqlService) GetBoyOne(id int64) (data *model.Boy, err error) {
	data, err = b.boy.GetOne(id)
	return
}

func (b *postgresqlService) CreateOne(req postgresqlRequest.Boy) (data *model.Boy, err error) {
	var bys []byte
	bys, err = json.Marshal(req.Informations)
	if err != nil {
		return nil, err
	}
	data, err = b.boy.CreateOne(&model.Boy{
		Name:         req.Name,
		Informations: model.Jsonb{RawMessage: json.RawMessage(bys)},
		Arr:          req.Arr,
	})
	return data, err
}
