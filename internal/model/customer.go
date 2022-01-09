package model

import "time"

const StaffNotDelCustomer = 0
const StaffDelCustomer = 1
const CustomerNotDelStaff = 0
const CustomerDelStaff = 1

type Customer struct {
	Id                  int    `gorm:"AUTO_INCREMENT;column:id;type:INT;primary_key" json:"id"`
	Corpid              string `gorm:"column:corpid;type:VARCHAR;size:255" json:"corpid"`
	ExternalUserid      string `gorm:"column:externalUserid;type:VARCHAR;size:255" json:"externalUserid"`
	ExternalUseridExtra string `gorm:"column:externalUseridExtra" json:"externalUseridExtra"`
	Name                string `gorm:"column:name;type:VARCHAR;size:300" json:"name"`
	OfficePosition      string `gorm:"column:officePosition;type:VARCHAR;size:128" json:"position"`
	Avatar              string `gorm:"column:avatar;type:VARCHAR;size:255" json:"avatar"`
	CorpName            string `gorm:"column:corpName;type:VARCHAR;size:64" json:"corpName"`
	CorpFullName        string `gorm:"column:corpFullName;type:VARCHAR;size:64" json:"corpFullName"`
	Type                int    `gorm:"column:type;type:int" json:"type"`
	Gender              int    `gorm:"column:gender;type:int" json:"Gender"`
	UnionId             string `gorm:"column:unionId;type:VARCHAR;size:64" json:"unionId"`
	UnionIdExtra        string `gorm:"column:unionIdExtra" json:"unionIdExtra"`
	ExternalProfile     string `gorm:"column:externalProfile;type:VARCHAR;size:500;" json:"externalProfile"`
	IsDelete            int    `gorm:"column:isDelete;type:int;default:0;" json:"isDelete"`
	CreateTime          int64  `gorm:"column:createTime;type:int;default:0;" json:"createTime"`
	UpdateTime          int64  `gorm:"column:updateTime;type:int;default:0;" json:"updateTime"`
	CustomerUri         string `gorm:"column:customerUri;type:VARCHAR" json:"customerUri"`
	CompanyId           int64  `gorm:"column:companyId" json:"companyId"`
	OpenId              string `gorm:"column:openid" json:"openid"`
}

type CustomerFollowerUserEs struct {
	Id                  int     `json:"id"`                    // 对应 qy_customer_follow_user 表 id
	UserId              string  `json:"user_id"`               // 对应 qy_customer_follow_user 表 userid
	StaffId             int     `json:"staff_id"`              // 对应 qy_customer_follow_user 表 staffId
	AddTime             uint    `json:"add_time"`              // 对应 qy_customer_follow_user 表 addTime
	AddChannel          int     `json:"add_channel"`           // 对应 qy_customer_follow_user 表 add_channal
	DelFollow           int     `json:"del_follow"`            // 对应 qy_customer_follow_user 表 delFollow
	DelExternal         int     `json:"del_external"`          // 对应 qy_customer_follow_user 表 delExternal
	Remark              string  `json:"remark"`                // 对应 qy_customer_follow_user 表 remark
	CustomerState       int     `json:"customer_state"`        // 根据 DelExternal 和 DelFollow 组合而成
	CustomerId          int     `json:"customer_id"`           // 对应 qy_customers 表 id
	Gender              int     `json:"gender"`                // 对应 qy_customers 表 gender
	CustomerName        string  `json:"customer_name"`         // 对应 qy_customers 表 name
	CustomerTags        []int   `json:"customer_tags"`         // 对应 qy_customer_tag 表 id
	LastTransactionTime uint    `json:"last_transaction_time"` // 对应 qy_customer_transaction 表 lastTransactionTime
	VolumeTotal         int     `json:"volume_total"`          // 对应 qy_customer_transaction 表 volumeTotal 总成交量，排序字段
	TransactionTotal    float64 `json:"transaction_total"`     // 对应 transactionTotal 表  总成交额，排序字段
	CompanyId           int64   `json:"company_id"`            // 对应 qy_customer_follow_user 表 CompanyId
}

func (Customer) TableName() string {
	return "qy_customers"
}

type (
	CustomerFollowUser struct {
		Id              int    `gorm:"AUTO_INCREMENT;column:id;type:INT;primary_key" json:"id"`
		CorpId          string `gorm:"column:corpid;type:VARCHAR;size:255" json:"corpId"`
		CompanyId       int64  `gorm:"column:companyId;type:INT;size:10" json:"companyId"`
		StaffId         int64  `gorm:"column:staffId;type:INT;size:10" json:"staffId"`
		CustomerId      int    `gorm:"column:customerId;type:INT;size:10" json:"customerId"`
		UserId          string `gorm:"column:userid;type:VARCHAR;size:255" json:"userId"`
		Remark          string `gorm:"column:remark;type:VARCHAR;size:255" json:"remark"`
		Description     string `gorm:"column:description;type:text" json:"description"`
		CreateTime      int64  `gorm:"column:createTime;type:int" json:"createTime"`
		TagNames        string `gorm:"column:tagNames;type:VARCHAR;size:255" json:"tagNames"`
		RemarkCorpName  string `gorm:"column:remarkCorpName;type:VARCHAR;size:255" json:"remarkCorpName"`
		RemarkMobiles   string `gorm:"column:remarkMobiles;type:VARCHAR;size:255" json:"remarkMobiles"`
		AddWay          uint16 `gorm:"column:addWay;type:int" json:"addWay"`
		Type            int    `gorm:"column:type;type:int" json:"type"`
		ExternalUserId  string `gorm:"column:externalUserId;type:VARCHAR;size:255" json:"externalUserId"`
		State           string `gorm:"column:state;type:VARCHAR;size:255" json:"state"`
		AddTime         uint64 `gorm:"column:addTime;type:int" json:"addTime"`
		AddChannal      int    `gorm:"column:add_channal;type:int" json:"addChannal"`
		DelExternal     int    `gorm:"column:delExternal;type:int" json:"delExternal"`
		DelExternalTime int64  `gorm:"column:delExternalTime;type:int" json:"delExternalTime"`
		DelFollow       int    `gorm:"column:delFollow;type:int" json:"delFollow"`
		DelFollowTime   int64  `gorm:"column:delFollowTime;type:int" json:"delFollowTime"`
	}
	CustomerFollowUserWith struct {
		CustomerFollowUser
		Customer            Customer            `gorm:"ForeignKey:CustomerId"`
		CustomerTags        []QyCustomerTag     `gorm:"ForeignKey:followUserId;references:id"`
		CustomerTransaction CustomerTransaction `gorm:"ForeignKey:customerId;references:customerId"`
	}

	CustomerFollowUser2 struct {
		Id        int    `gorm:"AUTO_INCREMENT;column:id;type:INT;primary_key" json:"id"`
		CorpId    string `gorm:"column:corpid;type:VARCHAR;size:255" json:"corpId"`
		CompanyId int64  `gorm:"column:companyId;type:INT;size:10" json:"companyId"`
		StaffId   int64  `gorm:"column:staffId;type:INT;size:10" json:"staffId"`
	}
)

func (CustomerFollowUser) TableName() string {
	return "qy_customer_follow_user"
}

type ExternalContactor struct {
	Id              int    `gorm:"AUTO_INCREMENT;column:id;type:INT;primary_key" json:"id"`
	CorpId          string `gorm:"column:corp_id" json:"corpId"`
	CompanyId       int64  `gorm:"column:company_id" json:"companyId"`
	CustomerId      int    `gorm:"column:customer_id" json:"customerId"`
	OpenId          string `gorm:"column:open_id" json:"openId"`
	ThirdExternalId string `gorm:"column:third_external_id" json:"thirdExternalId"`
	ThirdUnionId    string `gorm:"column:third_union_id" json:"thirdUnionId"`
	OwnExternalId   string `gorm:"column:own_external_id" json:"ownExternalId"`
	OwnUnionId      string `gorm:"column:own_union_id" json:"ownUnionId"`
	CreateTime      int64  `gorm:"column:create_time" json:"createTime"`
	IsDelete        uint8  `gorm:"column:is_delete" json:"isDelete"`
}

func (ExternalContactor) TableName() string {
	return "qy_external_contactor"
}

type QyCustomerTag struct {
	ID                int64     `gorm:"primaryKey;column:id" json:"id"`                                           // 自增ID
	CustomerID        int       `gorm:"index:idx_customer_id;default:0;column:customerId" json:"customerId"`      // 客户ID
	StaffID           int       `gorm:"index:idx_staff_id;default:0;column:staffId" json:"staffId"`               // 员工ID
	FollowUserID      int       `gorm:"index:idx_followUserId;default:0;column:followUserId" json:"followUserId"` // 成员与客户的关系id
	TagID             int       `gorm:"index:idx_tag_id;default:0;column:tagId" json:"tagId"`                     // 标签id
	QyCustomerTagType string    `gorm:"default:'';column:qy_customer_tag_type" json:"qyCustomerTagType"`          // 标签类型
	ModifyTime        time.Time `gorm:"default:CURRENT_TIMESTAMP;column:modifyTime" json:"modifyTime"`            // 更新时间
	CreateTime        int64     `gorm:"default:0;column:createTime" json:"createTime"`                            // 创建时间
}

func (QyCustomerTag) TableName() string {
	return "qy_customer_tag"
}

type CustomerTransaction struct {
	ID                        uint
	CompanyId                 uint       `gorm:"column:companyId"`
	ExternalUserId            string     `gorm:"column:externalUserId"`
	CustomerId                uint       `gorm:"column:customerId"`
	TransactionTotal          float64    `gorm:"column:transactionTotal"`
	VolumeTotal               int        `gorm:"column:volumeTotal"`
	LastTransactionTime       uint       `gorm:"column:lastTransactionTime"`
	LastTransactionPlatform   int        `gorm:"column:lastTransactionPlatform"`
	LastTransactionCorpShopId uint       `gorm:"column:lastTransactionCorpShopId"`
	CreateTime                uint       `gorm:"column:createTime"`
	UpdateTime                uint       `gorm:"column:updateTime"`
	ModifyTime                *time.Time `gorm:"column:modifyTime"`
}

func (CustomerTransaction) TableName() string {
	return "qy_customer_transaction"
}

type CustomerFollowUserList struct {
	Id             int64  `gorm:"column:id"`
	CompanyId      int    `gorm:"column:companyId"`
	Corpid         string `gorm:"column:corpid"`
	StaffId        int    `gorm:"column:staffId"`
	CustomerId     int    `gorm:"column:customerId"`
	ExternalUserid string `gorm:"column:externalUserid"`
	Userid         string `gorm:"column:userid"`
}

func (CustomerFollowUserList) TableName() string {
	return "qy_customer_follow_user"
}
