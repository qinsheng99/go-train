package mongoRequest

type RequestOne struct {
	DD
}

type RequestMany struct {
	Data []DD
}

type RequestFind struct {
	Name string `bson:"name" json:"name" form:"name"`
}

type DD struct {
	Name string `bson:"name" json:"name"`
	Age  int64  `bson:"age" json:"age"`
	Cve  string `bson:"cve" json:"cve"`
	Dep  string `bson:"dep" json:"dep"`
	Repo string `bson:"repo" json:"repo"`
}
