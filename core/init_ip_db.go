package core

import (
	"BlogServer/utlis/ipUtlis"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
	"strings"
)

var searcher *xdb.Searcher

func InitIPDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(dbPath) // 这里加载到内存了
	if err != nil {
		logrus.Fatalf("IP地址数据库加载失败 %s", err.Error())
		return
	}
	searcher = _searcher
}

func GetIPAddr(ip string) (addr string) {
	if ipUtlis.HasLocalIPAddr(ip) {
		return "内网"
	}
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("错误的IP地址 %s", err)
		return "异常地址"
	}
	// _addrList五个部分
	_addrList := strings.Split(region, "|")
	if len(_addrList) != 5 {
		// 不会有这个情况吧
		logrus.Warnf("异常的IP地址%s", ip)
		return "未知地址"
	}
	// 国家 0 省份 市 运营商
	country := _addrList[0]
	province := _addrList[2]
	city := _addrList[3]

	if province != "0" && city != "0" {
		return fmt.Sprintf("%s·%s", province, city)
	}
	if country != "0" && province != "0" {
		return fmt.Sprintf("%s·%s", country, province)
	}
	if country != "0" {
		return fmt.Sprintf("%s", country)
	}
	return region
}
