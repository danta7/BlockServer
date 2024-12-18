package global

import (
	"BlogServer/conf"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *conf.Config // 全局配置文件
	DB     *gorm.DB     // 数据库
	Redis  *redis.Client
)
