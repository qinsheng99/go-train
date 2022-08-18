package mongoRequest

type Request struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Name       string `bson:"name" json:"name"`
	Age        int64  `bson:"age" json:"age"`
	Cve        string `bson:"cve" json:"cve"`
	Dep        string `bson:"dep" json:"dep"`
	Repo       string `bson:"repo" json:"repo"`
}
