package response

type CeShiResponse struct {
	List     []interface{} `json:"list"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Total    int           `json:"total"`
	Ids      []int         `json:"ids"`
}

type ShardFailure struct {
	Index   string                 `json:"_index,omitempty"`
	Shard   int                    `json:"_shard,omitempty"`
	Node    string                 `json:"_node,omitempty"`
	Reason  map[string]interface{} `json:"reason,omitempty"`
	Status  string                 `json:"status,omitempty"`
	Primary bool                   `json:"primary,omitempty"`
}
