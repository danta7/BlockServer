package conf

import "fmt"

type QQ struct {
	AppID    string `yaml:"appID" json:"appID"`
	AppKey   string `yaml:"appKey" json:"appKey"`
	Redirect string `yaml:"redirect" json:"redirect"` // 回调地址
}

// Url 跳转地址
func (q QQ) Url() string {
	// 真正用的话这里要在qq互联中心配置审核
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
