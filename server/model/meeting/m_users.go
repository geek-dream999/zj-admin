// 自动生成模板MUsers
package meeting

import (
	"time"
)

// mUsers表 结构体  MUsers
type MUsers struct {
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:记录创建时间;"`  //记录创建时间 
    Email  string `json:"email" form:"email" gorm:"column:email;comment:用户电子邮件，需要是唯一的;size:255;"`  //用户电子邮件，需要是唯一的 
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:用户ID;size:20;"`  //用户ID 
    IsMuted  *bool `json:"isMuted" form:"isMuted" gorm:"column:is_muted;comment:是否禁言：0表示未禁言，1表示已禁言;"`  //是否禁言：0表示未禁言，1表示已禁言 
    Openid  string `json:"openid" form:"openid" gorm:"column:openid;comment:微信小程序返回的OpenID，用于标识微信小程序账号;size:255;"`  //微信小程序返回的OpenID，用于标识微信小程序账号 
    Password  string `json:"password" form:"password" gorm:"column:password;comment:用户密码，实际应用中建议使用密码哈希存储;size:255;"`  //用户密码，实际应用中建议使用密码哈希存储 
    Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号码，需要是唯一的;size:20;"`  //用户手机号码，需要是唯一的 
    Status  *bool `json:"status" form:"status" gorm:"column:status;comment:用户状态：1表示激活，0表示未激活;"`  //用户状态：1表示激活，0表示未激活 
    Unionid  string `json:"unionid" form:"unionid" gorm:"column:unionid;comment:微信小程序返回的UnionID，用于标识微信开放平台账号;size:255;"`  //微信小程序返回的UnionID，用于标识微信开放平台账号 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:记录更新时间;"`  //记录更新时间 
    Username  string `json:"username" form:"username" gorm:"column:username;comment:用户名，需要是唯一的;size:50;"`  //用户名，需要是唯一的 
}


// TableName mUsers表 MUsers自定义表名 m_users
func (MUsers) TableName() string {
    return "m_users"
}

