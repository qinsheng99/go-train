package model

import (
	"time"

	"github.com/qinsheng99/go-train/library/db"
)

type SbomPackage struct {
	Id                int64     `gorm:"column:id"`
	ParentId          int64     `gorm:"column:parent_id" description:"关联的父id"`
	PackageType       int8      `gorm:"column:package_type" description:"1:本身包名(仓库名)，2：内部引用依赖，3：外部依赖"`
	PackageName       string    `gorm:"column:package_name;type:varchar(128)" description:"包名"`
	Version           string    `gorm:"column:version;type:varchar(128)" description:"包名对应的版本"`
	OriginUrl         string    `gorm:"column:origin_url;type:varchar(128)" description:"上游url"`
	CpePackname       string    `gorm:"column:cpe_packname;type:varchar(128)" description:"cpe名称"`
	RepoName          string    `gorm:"column:repo_name;type:varchar(128)" description:"包对应的仓库名称"`
	OrganizationName  string    `gorm:"column:organization_name;type:varchar(128)" description:"对应的组织名称"`
	OrganizationId    int8      `gorm:"column:organization_id" description:"对应的组织id,默认值为1:1:来源src-openEuler;2:来源opengauss;3:来源mindspore;4:来源openLooKeng"`
	Releases          string    `gorm:"column:releases;type:varchar(512)" description:"包对应的发布版本"`
	ReleaseTime       string    `gorm:"column:release_time;type:varchar(32)" description:"当前版本发布时间"`
	Status            int8      `gorm:"column:status" description:"1:新增，2：更新，3：已删除"`
	PushStatus        int8      `gorm:"column:push_status" description:"0:未推送，1:已推送"`
	License           string    `gorm:"column:license;type:varchar(512)" description:"证书编号"`
	MainTainer        string    `gorm:"column:main_tainer;type:varchar(128)" description:"维护人"`
	MainTainLevel     int8      `gorm:"column:main_tain_level" description:"软件包维护级别"`
	Feature           string    `gorm:"column:feature;type:varchar(128)" description:"特性"`
	LatestVersion     string    `gorm:"column:latest_version;type:varchar(128)" description:"最新版本号"`
	LatestVersionTime string    `gorm:"column:latest_version_time;type:varchar(32)" description:"最新版本发布时间"`
	Summary           string    `gorm:"column:summary;type:text" description:"简介"`
	Decription        string    `gorm:"column:decription;type:text" description:"描述"`
	Owner             string    `gorm:"column:owner;type:varchar(128)" description:"包对应负责人"`
	Warehouse         string    `gorm:"column:warehouse;type:varchar(64)" description:"仓库"`
	CreateTime        time.Time `gorm:"column:create_time"`
	UpdateTime        time.Time `gorm:"column:update_time"`
}

func (s *SbomPackage) TableName() string {
	return "cve_sbom_package"
}

func (s *SbomPackage) Insert() (int64, error) {
	base := db.GetPostgresqlDb().Create(s)
	return base.RowsAffected, base.Error
}
