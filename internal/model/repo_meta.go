package model

import (
	"github.com/lib/pq"
)

type RepoMeta struct {
	Id                   int64          `json:"id,omitempty" gorm:"column:id;type:uuid"`
	Branch               string         `json:"branch,omitempty" gorm:"column:branch;type:text"`
	ProductType          string         `json:"product_type,omitempty" gorm:"column:product_type;type:text"`
	RepoName             string         `json:"repo_name,omitempty" gorm:"column:repo_name;type:text"`
	DownloadLocation     string         `json:"download_location,omitempty" gorm:"column:download_location;type:varchar(255)"`
	SpecDownloadUrl      string         `json:"spec_download_url,omitempty" gorm:"column:spec_download_url;type:varchar(255)"`
	PackageNames         pq.StringArray `gorm:"column:package_names;type:text[]" json:"package_names,omitempty"`
	PatchInfo            pq.StringArray `gorm:"column:patch_info;type:text[]" json:"patch_info,omitempty"`
	UpstreamDownloadUrls pq.StringArray `gorm:"column:upstream_download_urls;type:text[]" json:"upstream_download_urls,omitempty"`
}

func (r *RepoMeta) TableName() string {
	return "repo_meta"
}
