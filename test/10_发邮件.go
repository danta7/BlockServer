package main

import (
	"BlogServer/core"
	"BlogServer/flags"
	"BlogServer/global"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()

	em := global.Config.Email

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{"1052323212@qq.com"}
	e.Subject = "【蛋挞博客】账号注册"
	e.Text = []byte("你正在进行账号注册操作，这是你的验证码：4564,十分钟内有效")
	err := e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain))
	if err != nil && !strings.Contains(err.Error(), "short response:") {
		fmt.Println(err)
		return
	}
	fmt.Println("发送成功")
}
