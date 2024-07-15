package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// Register User register structure
type Register struct {
	Username     string `json:"userName" warehouse:"用户名"`
	Password     string `json:"passWord" warehouse:"密码"`
	NickName     string `json:"nickName" warehouse:"昵称"`
	HeaderImg    string `json:"headerImg" warehouse:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" warehouse:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" warehouse:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" warehouse:"[]uint 角色id"`
	Phone        string `json:"phone" warehouse:"电话号码"`
	Email        string `json:"email" warehouse:"电子邮箱"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string                `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
