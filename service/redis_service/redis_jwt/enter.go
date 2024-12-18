package redis_jwt

import (
	"BlogServer/global"
	"BlogServer/utlis/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type BlackType int8

const (
	UserBlackType   BlackType = 1 // 用户注销登录
	AdminBlackType  BlackType = 2 // 管理员操作下线
	DeviceBlackType BlackType = 3 // 其他设备把自己挤下来了
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
}

func (b BlackType) Msg() string {
	switch b {
	case UserBlackType:
		return "已注销"
	case AdminBlackType:
		return "禁止登录"
	case DeviceBlackType:
		return "设备下线"
	}
	return "已注销"
}

func ParseBlackType(val string) BlackType {
	switch val {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	}
	return UserBlackType
}

func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)
	claims, err := jwts.ParseToken(token)
	if err != nil || claims == nil {
		logrus.Errorf("token解析失败 %s", err)
		return
	}
	second := claims.ExpiresAt - time.Now().Unix()
	_, err = global.Redis.Set(key, value.String(), time.Duration(second)*time.Second).Result()
	if err != nil {
		logrus.Errorf("redis添加黑名单失败 %s", err)
		return
	}

}

func HasTokenBlack(token string) (blkType BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(key).Result()
	if err != nil {
		return
	}
	blkType = ParseBlackType(value)
	return blkType, true
}

func HasTokenBlackByGin(c *gin.Context) (blkType BlackType, ok bool) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
	}
	return HasTokenBlack(token)
}
