package elasticsearch

import (
	"github.com/olivere/elastic/v7"
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

func GetES() (*ES, error) {
	var c = Config{
		URL:      "http://localhost:9200",
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
