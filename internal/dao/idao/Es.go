package idao

import (
	"github.com/qinsheng99/goWeb/api/entity/ceshi/request"
	crequest "github.com/qinsheng99/goWeb/api/entity/ceshi/request"
	"github.com/qinsheng99/goWeb/api/entity/ceshi/response"
	"github.com/qinsheng99/goWeb/internal/model"
	"github.com/qinsheng99/goWeb/library/funcTest"

	"github.com/olivere/elastic/v7"
)

type EsImp interface {
	PostEsList(q request.CeShiRequest) (r response.CeShiResponse, err error)
	GetEsList(q request.CeShiGetRequest) (r response.CeShiResponse, err error)
	Delete(id int) (*elastic.DeleteResponse, error)
	GetEsById(q request.CeShiRequest) ([]model.CustomerFollowerUserEs, error)
	DeleteEs() error
	InsertElastic(list []*model.CustomerFollowerUserEs)
	Get(q crequest.CeShiGetRequest) (*elastic.SearchResult, error)
}

func PQuery(q request.CeShiRequest) *elastic.BoolQuery {
	b := elastic.NewBoolQuery()

	if q.StaffId > 0 {
		b.Must(elastic.NewTermQuery("staff_id", q.StaffId))
	}

	if q.AddTime > 0 {
		b.Must(elastic.NewTermQuery("add_time", q.AddTime))
	}

	if q.UserId != "" {
		b.Must(elastic.NewMatchQuery("user_id", q.UserId))
	}

	if q.CompanyId > 0 {
		b.Must(elastic.NewMatchQuery("company_id", q.CompanyId))
	}

	adds := funcTest.FilterIntSlice(q.AddChannel)
	if len(adds) > 0 {
		b.Must(elastic.NewTermsQuery("add_channel", funcTest.IntSliceToInf(adds)...))
	}
	return b
}

func GQuery(q request.CeShiGetRequest) *elastic.BoolQuery {
	boolq := elastic.NewBoolQuery()

	addIds := funcTest.FilterIntSlice(q.AddChannel)

	if len(addIds) > 0 {
		boolq.Must(elastic.NewTermsQuery("add_channel", funcTest.IntSliceToInf(addIds)...))
	}

	if q.CompanyId > 0 {
		boolq.Must(elastic.NewTermQuery("company_id", q.CompanyId))
	}

	if len(q.Tags) > 0 && q.Option != "" {
		b := elastic.NewBoolQuery()
		for _, v := range q.Tags {
			if q.Option == "and" {
				b.Must(elastic.NewTermQuery("customer_tags", v))
			} else {
				b.Should(elastic.NewTermQuery("customer_tags", v))
			}
		}
		boolq.Must(b)
	}

	if q.Id > 0 {
		boolq.Must(elastic.NewTermQuery("id", q.Id))
	}

	if q.LossState > 0 {
		boolq.Must(elastic.NewTermQuery("customer_state", q.LossState))
	}

	return boolq
}
