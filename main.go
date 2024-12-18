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
	core.InitLogrus()         // 初始化记录日志
	global.DB = core.InitDB() // 初始化数据库 包括读写分类等
	global.Redis = core.InitRedis()

	flags.Run()

	// 启动web程序
	router.Run()
}
