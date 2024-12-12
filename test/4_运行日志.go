package main

import (
	"BlogServer/core"
	"BlogServer/flags"
	"BlogServer/global"
	"BlogServer/service/log_service"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	log := log_service.NewRuntimeLog("同步文章数据", log_service.RuntimeDataHour)
	log.SetItem("文章1", 11)
	log.Save()

	log.SetItem("文章2", 22)
	log.Save()
}
