package main

import (
	"BlockServer/core"
	"BlockServer/flags"
	"BlockServer/global"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
}
