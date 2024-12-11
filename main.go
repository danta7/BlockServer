package main

import (
	"BlogServer/core"
	"BlogServer/flags"
	"BlogServer/global"
	"BlogServer/router"
)

func main() {
	flags.Parse() // 绑定命令行参数
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	flags.Run()

	// 启动web程序
	router.Run()
}
