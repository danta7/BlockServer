package global

import (
	"BlockServer/conf"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
)
