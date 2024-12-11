package core

import (
	"BlogServer/conf"
	"BlogServer/flags"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadConf() (c *conf.Config) {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}
	c = new(conf.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("yaml 配置文件格式错误:%v", err))
	}

	fmt.Printf("读取配置文件 %s 成功\n ", flags.FlagOptions.File)
	return
}
