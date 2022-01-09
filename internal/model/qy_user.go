package model

import "time"

type QyUser struct {
	Id               int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                         // 自增ID
	StaffId          int       `gorm:"column:staffId;default:0;NOT NULL" json:"staffId"`                       // 员工 ID
	Corpid           string    `gorm:"column:corpid;NOT NULL" json:"corpid"`                                   // 授权ID
	Source           int       `gorm:"column:source;default:0;NOT NULL" json:"source"`                         // 1手机号，2企业微信，3钉钉
	Sort             string    `gorm:"column:sort;NOT NULL" json:"sort"`                                       // 部门内的排序值
	Userid           string    `gorm:"column:userid;NOT NULL" json:"userid"`                                   // 第三方用户ID
	Gender           int       `gorm:"column:gender;default:0;NOT NULL" json:"gender"`                         // 0表示未定义，1表示男性，2表示女性
	Unionid          string    `gorm:"column:unionid;NOT NULL" json:"unionid"`                                 // 唯一标识
	OpenUserid       string    `gorm:"column:openUserid;NOT NULL" json:"openUserid"`                           // 成员的open_userid，全局唯一
	Name             string    `gorm:"column:name;NOT NULL" json:"name"`                                       // 名字
	EnglishName      string    `gorm:"column:englishName;NOT NULL" json:"englishName"`                         // 英文名
	Avatar           string    `gorm:"column:avatar;NOT NULL" json:"avatar"`                                   // 头像
	ThumbAvatar      string    `gorm:"column:thumbAvatar;NOT NULL" json:"thumbAvatar"`                         // 头像缩略图url
	Mobile           string    `gorm:"column:mobile;NOT NULL" json:"mobile"`                                   // 手机号
	Email            string    `gorm:"column:email;NOT NULL" json:"email"`                                     // 邮箱
	OfficePosition   string    `gorm:"column:officePosition;NOT NULL" json:"officePosition"`                   // 职位
	Department       string    `gorm:"column:department;NOT NULL" json:"department"`                           // 部门
	MainDepartment   string    `gorm:"column:mainDepartment;NOT NULL" json:"mainDepartment"`                   // 主部门
	JobNumber        string    `gorm:"column:jobNumber;NOT NULL" json:"jobNumber"`                             // 工号
	HiredDate        int       `gorm:"column:hiredDate;default:0;NOT NULL" json:"hiredDate"`                   // 入职时间
	Birthday         int       `gorm:"column:birthday;default:0;NOT NULL" json:"birthday"`                     // 生日
	IsActive         int       `gorm:"column:isActive;default:0;NOT NULL" json:"isActive"`                     // 是否激活
	IsAdmin          int       `gorm:"column:isAdmin;default:0;NOT NULL" json:"isAdmin"`                       // 是否管理员
	HideMobile       int       `gorm:"column:hideMobile;default:0;NOT NULL" json:"hideMobile"`                 // 是否隐藏手机号
	State            int       `gorm:"column:state;default:0;NOT NULL" json:"state"`                           // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。
	IsLeader         int       `gorm:"column:isLeader;default:0;NOT NULL" json:"isLeader"`                     // 是否主管
	IsLeaderInDept   string    `gorm:"column:isLeaderInDept;NOT NULL" json:"isLeaderInDept"`                   // 表示在所在的部门内是否为上级
	Telephone        string    `gorm:"column:telephone;NOT NULL" json:"telephone"`                             // 座机
	Alias            string    `gorm:"column:alias;NOT NULL" json:"alias"`                                     // 别名
	Extattr          string    `gorm:"column:extattr" json:"extattr"`                                          // 扩展属性
	Address          string    `gorm:"column:address;NOT NULL" json:"address"`                                 // 地址
	ExternalProfile  string    `gorm:"column:externalProfile;NOT NULL" json:"externalProfile"`                 // 成员对外属性
	ExternalPosition string    `gorm:"column:externalPosition;NOT NULL" json:"externalPosition"`               // 对外职务
	QrCode           string    `gorm:"column:qrCode;NOT NULL" json:"qrCode"`                                   // 员工二维码
	InfoJson         string    `gorm:"column:infoJson" json:"infoJson"`                                        // 其他信息
	ArchiveMsgState  int       `gorm:"column:archive_msg_state;default:0;NOT NULL" json:"archive_msg_state"`   // 消息存档开通状态: 0未开通，1已开通
	ModifyTime       time.Time `gorm:"column:modifyTime;default:CURRENT_TIMESTAMP;NOT NULL" json:"modifyTime"` // 更新时间
	OwnAuthorized    int       `gorm:"column:ownAuthorized;default:0;NOT NULL" json:"ownAuthorized"`           // 自建应用是否授权0否1已授权
	ThirdAuthorized  int       `gorm:"column:thirdAuthorized;default:1;NOT NULL" json:"thirdAuthorized"`       // 第三方应用是否授权0否1已授权
	IsAuthUser       int       `gorm:"column:isAuthUser;default:0;NOT NULL" json:"isAuthUser"`                 // 是否企业授权人，0否 1 是
	CreateTime       int       `gorm:"column:createTime;default:0;NOT NULL" json:"createTime"`                 // 创建时间
	IsDelete         int       `gorm:"column:isDelete;default:0;NOT NULL" json:"isDelete"`                     // 0 未删除 1 已删除
	DeleteTime       int       `gorm:"column:deleteTime;default:0;NOT NULL" json:"deleteTime"`                 // 删除时间
	OpenUserId       string    `gorm:"column:open_user_id;NOT NULL" json:"open_user_id"`                       // 成员在企微的全局唯一标识
}

type QyUser1 struct {
	Id     int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name   string `gorm:"column:name;NOT NULL" json:"name"` // 名字
	Avatar string `gorm:"column:avatar;NOT NULL" json:"avatar"`
}

func (QyUser) TableName() string {
	return "qy_user"
}
