package main

import (
	"BlogServer/core"
	"BlogServer/flags"
	"BlogServer/global"
	"BlogServer/utlis/jwts"
	"fmt"
)

func main() {
	flags.Parse() // 绑定命令行参数
	global.Config = core.ReadConf()
	core.InitLogrus()
	token, err := jwts.GetToken(jwts.Claims{
		UserID: 2,
		Role:   1,
	})
	fmt.Println(token, err)
	//cls, err := jwts.ParseToken(token)
	//fmt.Println(cls, err)
}
