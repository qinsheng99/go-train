package elasticsearch

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/qinsheng99/go-train/library/config"
)

type (
	ES struct {
		*elastic.Client
	}

	Config struct {
		URL      string `json:"url"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ConfLoader func(v interface{}) error
)

const (
	ESEtcdConfIndex = "watch.elasticsearch"

	ESHighlightPreTag  = "<span>"
	ESHighlightPostTag = "</span>"
)

func GetES(cfg *config.EsConfig) (*ES, error) {
	var c = Config{
		URL:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "",
		Username: "",
	}
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(c.URL),
		elastic.SetSniff(false),
	}

	if c.Password != "" {
		options = append(options, elastic.SetBasicAuth(c.Username, c.Password))
	}

	client, err := elastic.NewClient(options...)
	if err != nil {
		return nil, err
	}

	return &ES{
		Client: client,
	}, nil
}
