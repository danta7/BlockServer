package conf

type Email struct {
	Domain       string `json:"domain" yaml:"domain"`
	Port         int    `json:"port" yaml:"port"`
	SendEmail    string `json:"sendEmail" yaml:"sendEmail"`
	AuthCode     string `json:"authCode" yaml:"authCode"` // 授权码 验证码
	SendNickname string `json:"sendNickname" yaml:"sendNickname"`
	SSL          bool   `json:"SSL" yaml:"ssl"`
	TLS          bool   `json:"TLS" yaml:"tls"`
}
