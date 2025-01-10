package global

import (
	"BlogServer/conf"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
	"sync"
)

const Version = "10.0.1"

var (
	Config           *conf.Config // 全局配置文件
	DB               *gorm.DB     // 数据库
	Redis            *redis.Client
	CaptchaStore     = base64Captcha.DefaultMemStore // 用于存储验证码ID和对应的验证码文本
	EmailVerifyStore = sync.Map{}
)
