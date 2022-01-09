package model

import "time"

type QyTag struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                         // 自增ID
	Type       int       `gorm:"column:type;default:0;NOT NULL" json:"type"`                             // 标签类型 0 平台标签 1 企业微信标签
	GroupId    int       `gorm:"column:groupId;default:0;NOT NULL" json:"groupId"`                       // 标签组id
	Name       string    `gorm:"column:name;NOT NULL" json:"name"`                                       // 字标签name
	ExtId      string    `gorm:"column:extId;NOT NULL" json:"extId"`                                     // 企业微信标签id
	ExtGroupId string    `gorm:"column:extGroupId;NOT NULL" json:"extGroupId"`                           // 企业微信标签组id
	CorpId     string    `gorm:"column:corpId;NOT NULL" json:"corpId"`                                   // 企业主键id
	CreateTime int       `gorm:"column:createTime;default:0;NOT NULL" json:"createTime"`                 // 入库时间
	CreatedAt  int       `gorm:"column:createdAt;default:0;NOT NULL" json:"createdAt"`                   // 标签创建时间
	Sort       int       `gorm:"column:sort;default:0;NOT NULL" json:"sort"`                             // 排序
	IsDelete   int       `gorm:"column:isDelete;default:0;NOT NULL" json:"isDelete"`                     // 是否删除
	ModifyTime time.Time `gorm:"column:modifyTime;default:CURRENT_TIMESTAMP;NOT NULL" json:"modifyTime"` // 更新时间
}

type QyTag1 struct {
	Id    int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Type  int    `gorm:"column:type;default:0;NOT NULL" json:"type"`
	ExtId string `gorm:"column:extId;NOT NULL" json:"extId"`
	Sort  int    `gorm:"column:sort;default:0;NOT NULL" json:"sort"`
	Name  string `gorm:"column:name;NOT NULL" json:"name"`
}

func (QyTag) TableName() string {
	return "qy_tag"
}
