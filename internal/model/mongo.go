package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id" json:"-"`
	Name string             `bson:"name" json:"name"`
	Age  int64              `bson:"age" json:"age"`
	Cve  string             `bson:"cve" json:"cve"`
	Dep  string             `bson:"dep" json:"dep"`
	Repo string             `bson:"repo" json:"repo"`
}

type DCLA struct {
	URL          string `bson:"url" json:"url" required:"true"`
	Text         string `bson:"text" json:"text" required:"true"`
	OrgSignature []byte `bson:"org_signature" json:"-"`

	DCLAInfo `bson:",inline"`
}
type DCLAInfo struct {
	Fields           []DField `bson:"fields" json:"fields,omitempty"`
	Language         string   `bson:"lang" json:"lang" required:"true"`
	CLAHash          string   `bson:"cla_hash" json:"cla_hash" required:"true"`
	OrgSignatureHash string   `bson:"signature_hash" json:"signature_hash,omitempty"`
}

type DField struct {
	ID          string `bson:"id" json:"id" required:"true"`
	Title       string `bson:"title" json:"title" required:"true"`
	Type        string `bson:"type" json:"type" required:"true"`
	Description string `bson:"desc" json:"desc,omitempty"`
	Required    bool   `bson:"required" json:"required"`
}

type DWuKong struct {
	Id       string        `bson:"id"        json:"id"`
	Samples  []DSample     `bson:"samples"   json:"samples"`
	Pictures []DictureInfo `bson:"pictures"  json:"pictures"`
}

type DSample struct {
	Num  int    `bson:"num"              json:"num"`
	Name string `bson:"name"             json:"name"`
}

type DictureInfo struct {
	Desc  string `bson:"desc"            json:"desc"`
	Link  string `bson:"link"            json:"link"`
	Style string `bson:"style"           json:"style"`
}
