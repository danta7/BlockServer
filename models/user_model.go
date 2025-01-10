package models

import (
	"BlogServer/models/enum"
	"time"
)

// UserModel 用户信息表
type UserModel struct {
	Model
	Username       string                  `gorm:"size:32" json:"username"`
	Nickname       string                  `gorm:"size:32" json:"nickname"`
	Avatar         string                  `gorm:"size:256" json:"avatar"` // 头像
	Abstract       string                  `gorm:"size:256" json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"` // 注册来源
	CodeAge        int8                    `json:"codeAge"`        // 码龄
	Password       string                  `gorm:"size:64" json:"-"`
	Email          string                  `gorm:"size:256" json:"email"`
	OpenID         string                  `gorm:"size:64" json:"openID"` // 第三方登录的唯一id
	Role           enum.RoleType           `json:"role"`                  // 角色 1：管理员 2：普通用户 3：访客
}

// UserConfModel 用户配置表
type UserConfModel struct {
	UserID             uint       `gorm:"unique" json:"userID"`
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`                    // 外键：引用UserModel的主键
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"` // 兴趣标签
	UpdateUsernameDate *time.Time `json:"updateUsernameDate"`                            // 上次修改用户名的时间 指针：可以为nil
	OpenCollect        bool       `json:"openCollect"`                                   // 是否公开我的收藏
	OpenFans           bool       `json:"openFans"`                                      // 是否公开我的粉丝
	OpenFollow         bool       `json:"openFollow"`                                    // 是否公开我的关注
	HomeStyleID        uint       `json:"homeStyleID"`                                   // 主页样式ID
}
