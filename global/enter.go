package global

import (
	"BlogServer/conf"
	"gorm.io/gorm"
)

var (
	Config *conf.Config // 全局配置文件
	DB     *gorm.DB     // 数据库
)
