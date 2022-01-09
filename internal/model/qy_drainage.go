package model

import "time"

type QyDrainage struct {
	Id           int64                 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	CompanyId    int64                 `gorm:"column:company_id;default:0;NOT NULL" json:"company_id"`                   // 公司ID
	StaffId      int64                 `gorm:"column:staff_id;default:0;NOT NULL" json:"staff_id"`                       // 创建人
	LinkName     string                `gorm:"column:link_name;NOT NULL" json:"link_name"`                               // 链接名称
	Uri          string                `gorm:"column:uri;NOT NULL" json:"uri"`                                           // uri
	QrCodeId     int                   `gorm:"column:qr_code_id;default:0;NOT NULL" json:"qr_code_id"`                   // qy_contact_way表id
	BackType     int                   `gorm:"column:back_type;default:0;NOT NULL" json:"back_type"`                     // 1为默认，2为无主题，3为自定义
	BackImage    string                `gorm:"column:back_image;NOT NULL" json:"back_image"`                             // 背景图
	BrandName    string                `gorm:"column:brand_name;NOT NULL" json:"brand_name"`                             // 品牌名
	BrandLogo    string                `gorm:"column:brand_logo;NOT NULL" json:"brand_logo"`                             // 品牌logo
	BrowseNum    int                   `gorm:"column:browse_num;default:0;NOT NULL" json:"browse_num"`                   // 浏览数
	AddNum       int                   `gorm:"column:add_num;default:0;NOT NULL" json:"add_num"`                         // 添加数，对应type为1使用
	UserIds      string                `gorm:"column:user_ids;NOT NULL" json:"user_ids"`                                 // 客户经理成员，对应type为1使用
	TagIds       string                `gorm:"column:tag_ids;NOT NULL" json:"tag_ids"`                                   // 标签id组，对应type为1使用
	ShortLink    string                `gorm:"column:short_link;NOT NULL" json:"short_link"`                             // 短链接
	IsDelete     int                   `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete"`                     // 软删除，0为删除，1删除
	Type         int                   `gorm:"column:type;default:0;NOT NULL" json:"type"`                               // 引流类型，1为加人引流，2为加群引流
	ServiceUsers string                `gorm:"column:service_users;NOT NULL" json:"service_users"`                       // 客服，对应type为2使用
	CreateTime   int64                 `gorm:"column:create_time;default:0;NOT NULL" json:"create_time"`                 // 创建时间
	UpdateTime   int64                 `gorm:"column:update_time;default:0;NOT NULL" json:"update_time"`                 // 更新时间
	DeleteTime   int64                 `gorm:"column:delete_time;default:0;NOT NULL" json:"delete_time"`                 // 删除时间
	ModifyTime   time.Time             `gorm:"column:modify_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"modify_time"` // 维护字段
	Staff        QyStaff               `gorm:"foreignKey:StaffId"`
	Code         QyContactWay          `gorm:"foreignKey:QrCodeId"`
	Group        []QyDrainageLinkGroup `gorm:"foreignKey:LinkUri;references:Uri"`
}

func (QyDrainage) TableName() string {
	return "qy_drainage"
}

type QyStaff struct {
	Id         int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                         // 自增ID
	Phone      string    `gorm:"column:phone;NOT NULL" json:"phone"`                                     // 手机号
	CompanyId  int       `gorm:"column:companyId;default:0;NOT NULL" json:"companyId"`                   // 公司 ID
	UserInfoId int       `gorm:"column:userInfoId;default:0;NOT NULL" json:"userInfoId"`                 // 用户中台 userinfoId
	StaffUri   string    `gorm:"column:staffUri;NOT NULL" json:"staffUri"`                               // 用户中台 uri
	Name       string    `gorm:"column:name;NOT NULL" json:"name"`                                       // 名字
	Avatar     string    `gorm:"column:avatar;NOT NULL" json:"avatar"`                                   // 头像
	Pwd        string    `gorm:"column:pwd;NOT NULL" json:"pwd"`                                         // 密码
	IsAuthUser int       `gorm:"column:isAuthUser;default:0;NOT NULL" json:"isAuthUser"`                 // 是否企业授权人，0否 1 是
	IsBlock    int       `gorm:"column:isBlock;default:0;NOT NULL" json:"isBlock"`                       // 是否停用 0 否 1 停用
	CreateTime int       `gorm:"column:createTime;default:0;NOT NULL" json:"createTime"`                 // 创建时间
	ModifyTime time.Time `gorm:"column:modifyTime;default:CURRENT_TIMESTAMP;NOT NULL" json:"modifyTime"` // 更新时间
	IsDelete   int       `gorm:"column:isDelete;default:0;NOT NULL" json:"isDelete"`                     // 软删除
	IsBind     int       `gorm:"column:is_bind;default:0;NOT NULL" json:"is_bind"`                       // 绑定账号标记 0: 未绑定 1:已绑定
}

func (QyStaff) TableName() string {
	return "qy_staff"
}

type QyContactWay struct {
	Id             int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	CompanyId      int64     `gorm:"column:company_id;default:0;NOT NULL" json:"company_id"`                   // 公司ID
	StaffId        int64     `gorm:"column:staff_id;default:0;NOT NULL" json:"staff_id"`                       // 创建人
	ConfigId       string    `gorm:"column:config_id;NOT NULL" json:"config_id"`                               // 企微config_id
	PushMsgId      int64     `gorm:"column:push_msg_id;default:0;NOT NULL" json:"push_msg_id"`                 // 推送消息ID
	State          string    `gorm:"column:state;NOT NULL" json:"state"`                                       // 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符
	QrCodeType     int       `gorm:"column:qr_code_type;default:0;NOT NULL" json:"qr_code_type"`               // 活码类型 0 普通 1 订阅活码  2 活动生成邀请活码 3 活动生成邀请活码 4 云商城客服活码 5 渠道注册码
	ContactWayName string    `gorm:"column:contact_way_name;NOT NULL" json:"contact_way_name"`                 // 名称
	Type           int       `gorm:"column:type;default:0;NOT NULL" json:"type"`                               // 联系方式类型,1-单人, 2-多人
	Scene          int       `gorm:"column:scene;default:0;NOT NULL" json:"scene"`                             // 场景，1-在小程序中联系，2-通过二维码联系
	Style          int       `gorm:"column:style;default:0;NOT NULL" json:"style"`                             // 在小程序中联系时使用的控件样式
	Remark         string    `gorm:"column:remark;NOT NULL" json:"remark"`                                     // 联系方式的备注信息，用于助记，不超过30个字符
	SkipVerify     int       `gorm:"column:skip_verify;default:1;NOT NULL" json:"skip_verify"`                 // 外部客户添加时是否无需验证，1为true 0为false
	QrCode         string    `gorm:"column:qr_code;NOT NULL" json:"qr_code"`                                   // 二维码地址
	Logo           string    `gorm:"column:logo;NOT NULL" json:"logo"`                                         // logo
	DiyQrCode      string    `gorm:"column:diy_qr_code;NOT NULL" json:"diy_qr_code"`                           // 自定义图片的二维码URL
	AddTimes       int       `gorm:"column:add_times;default:0;NOT NULL" json:"add_times"`                     // 添加次数
	IsDelete       int       `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete"`                     // 删除状态
	CreateTime     int       `gorm:"column:create_time;default:0;NOT NULL" json:"create_time"`                 // 创建时间
	ModifyTime     time.Time `gorm:"column:modify_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"modify_time"` // 更新时间
	DeleteTime     int       `gorm:"column:delete_time;default:0;NOT NULL" json:"delete_time"`                 // 删除时间
}

func (QyContactWay) TableName() string {
	return "qy_contact_way"
}

type QyDrainageLinkGroup struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                           // 自增ID
	LinkUri     string    `gorm:"column:link_uri;NOT NULL" json:"link_uri"`                                 // 引流链接uri
	GroupChatId string    `gorm:"column:group_chat_id;NOT NULL" json:"group_chat_id"`                       // 客户群ID
	QrcodeUrl   string    `gorm:"column:qrcode_url;NOT NULL" json:"qrcode_url"`                             // 群二维码
	IsDelete    int       `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete"`                     // 0 未删除1已删除
	CreateTime  int       `gorm:"column:create_time;default:0;NOT NULL" json:"create_time"`                 // 创建时间
	ModifyTime  time.Time `gorm:"column:modify_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"modify_time"` // 更新时间
}

func (QyDrainageLinkGroup) TableName() string {
	return "qy_drainage_link_group"
}
