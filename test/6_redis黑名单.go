package main

import (
	"BlogServer/core"
	"BlogServer/flags"
	"BlogServer/global"
	"BlogServer/service/redis_service/redis_jwt"
	"fmt"
)

func main() {
	flags.Parse() // 绑定命令行参数
	global.Config = core.ReadConf()
	core.InitLogrus() // 初始化记录日志
	global.Redis = core.InitRedis()

	//token, err := jwts.GetToken(jwts.Claims{
	//	UserID: 2,
	//	Role:   1,
	//})
	//fmt.Println(token, err)
	//redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	//blk, ok := redis_jwt.HasTokenBlack(token)
	//fmt.Println(blk, ok)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjIsInVzZXJuYW1lIjoiIiwicm9sZSI6MSwiZXhwIjoxNzM0NTAwNzYwLCJpc3MiOiJkYW50YSJ9.OdH518pkBj5hBb8NV6WqoJdQTWpST3d0P_Pu5JI4gaY"
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}
