package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/qinsheng99/go-train/api/entity/ceshi/request"
	crequest "github.com/qinsheng99/go-train/api/entity/ceshi/request"
	"github.com/qinsheng99/go-train/api/entity/ceshi/response"
	Err "github.com/qinsheng99/go-train/err"
	"github.com/qinsheng99/go-train/internal/dao/idao"
	"github.com/qinsheng99/go-train/internal/model"
	"github.com/qinsheng99/go-train/library/elasticsearch"

	"github.com/olivere/elastic/v7"
)

const (
	EsIndexCustomerFollowerUserListIndex = "customer_follower_user_list_index"
	EsIndexCeshi                         = "ceshi_index"
)

type EsDao struct {
	Es *elasticsearch.ES
}

func NewEsDao(es *elasticsearch.ES) idao.EsImp {
	return &EsDao{
		Es: es,
	}
}

func (e *EsDao) PostEsList(q request.CeShiRequest) (r response.CeShiResponse, err error) {
	var data model.CustomerFollowerUserEs
	page := q.Page
	pageSize := q.PageSize
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}
	from := (page - 1) * pageSize
	query := e.Es.Client.Search().
		Index(EsIndexCustomerFollowerUserListIndex).
		TrackScores(true).
		Query(idao.PQuery(q)).
		From(from).
		Size(pageSize)

	query.Sort("add_time", false)
	res, err := query.Do(context.Background())

	if err != nil {
		return r, err
	}
	r = response.CeShiResponse{
		Total:    int(res.Hits.TotalHits.Value),
		Page:     q.Page,
		PageSize: q.PageSize,
	}

	for _, v := range res.Hits.Hits {
		err := json.Unmarshal(v.Source, &data)
		if err != nil {
			return response.CeShiResponse{}, err
		}
		r.List = append(r.List, data)
	}
	return r, nil
}

func (e *EsDao) Delete(id int) (*elastic.DeleteResponse, error) {
	res, err := e.Es.Client.Delete().Index(EsIndexCustomerFollowerUserListIndex).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e *EsDao) GetEsById(q request.CeShiRequest) (r []model.CustomerFollowerUserEs, err error) {
	var data model.CustomerFollowerUserEs
	if len(q.Ids) <= 0 {
		return r, Err.Mesage("缺少id或id不存在")
	}
	for _, v := range q.Ids {
		res, err := e.Es.Client.Get().
			Index(EsIndexCustomerFollowerUserListIndex).Id(strconv.Itoa(int(v))).Do(context.Background())
		if err != nil {
			if elastic.IsNotFound(err) {
				continue
			}
			return r, err
		}
		_ = json.Unmarshal(res.Source, &data)
		r = append(r, data)
	}
	if len(r) == 0 {
		msg := fmt.Sprintf("data is empty and id in %v", q.Ids)
		return nil, Err.Mesage(msg)
	}

	return r, nil
}

func (e *EsDao) GetEsList(q request.CeShiGetRequest) (response.CeShiResponse, error) {
	var data model.CustomerFollowerUserEs
	var r response.CeShiResponse
	page := q.Page
	pageSize := q.PageSize
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}
	from := (page - 1) * pageSize
	res, err := e.Es.Client.Search().
		Index(EsIndexCustomerFollowerUserListIndex).
		Query(idao.GQuery(q)).
		From(from).
		Size(pageSize).Sort("add_time", false).
		Do(context.Background())

	if err != nil {
		if elastic.IsNotFound(err) {
			return response.CeShiResponse{}, nil
		}
		return response.CeShiResponse{}, err
	}
	r = response.CeShiResponse{
		Total:    int(res.Hits.TotalHits.Value),
		PageSize: pageSize,
		Page:     page,
	}

	for _, v := range res.Hits.Hits {
		err := json.Unmarshal(v.Source, &data)
		if err != nil {
			return r, err
		}
		r.List = append(r.List, data)
		r.Ids = append(r.Ids, data.Id)
	}
	return r, nil
}

func (e *EsDao) DeleteEs() error {
	if _, err := e.Es.Client.DeleteIndex(EsIndexCustomerFollowerUserListIndex).Do(context.Background()); err != nil {
		if elastic.IsNotFound(err) {
			return nil
		}
		return err
	}
	return nil
}

func (e *EsDao) InsertElastic(list []*model.CustomerFollowerUserEs) {
	for _, customerFollowerUser := range list {
		_, err := e.Es.Client.Index().
			Index(EsIndexCustomerFollowerUserListIndex).
			Id(strconv.Itoa(customerFollowerUser.Id)).
			BodyJson(customerFollowerUser).
			Do(context.Background())
		if err != nil {
			continue
		}
	}
}

func (e *EsDao) Get(q crequest.CeShiGetRequest) (*elastic.SearchResult, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("company_id", q.CompanyId))
	res, err := e.Es.Search().Index(EsIndexCustomerFollowerUserListIndex).Query(query).Do(context.Background())
	if err != nil {
		if elastic.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (e *EsDao) RefreshCeshi(datas []*model.CeshiEs) {
	for _, v := range datas {
		_, err := e.Es.Client.Index().
			Index(EsIndexCeshi).
			Id(strconv.Itoa(v.Id)).
			BodyJson(v).
			Do(context.Background())
		if err != nil {
			continue
		}
	}
}

func (e *EsDao) GetAllEsData() (datas []model.CeshiEs, err error) {
	var data model.CeshiEs
	res, err := e.Es.Client.Search().Index(EsIndexCeshi).
		Sort("id", true).
		Size(20).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	for _, v := range res.Hits.Hits {
		_ = json.Unmarshal(v.Source, &data)
		datas = append(datas, data)
	}
	return
}
