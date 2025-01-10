package email_service

import (
	"BlogServer/global"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

// SendRegisterCode 注册账号
func SendRegisterCode(to string, code string) error {
	em := global.Config.Email
	text := fmt.Sprintf("你正在进行账号注册操作，这是你的验证码:%s,十分钟内有效", code)
	subject := fmt.Sprintf("【%s】账号注册 ", em.SendNickname)
	return sendEmail(to, subject, text)
}

// SendResetPwdCode 重置密码
func SendResetPwdCode(to string, code string) error {
	em := global.Config.Email
	text := fmt.Sprintf("你正在进行密码重置操作，这是你的验证码:%s,十分钟内有效", code)
	subject := fmt.Sprintf("【%s】重置密码 ", em.SendNickname)
	return sendEmail(to, subject, text)
}

func sendEmail(to, subject, text string) (err error) {
	em := global.Config.Email

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	err1 := e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain))
	if err1 != nil && !strings.Contains(err1.Error(), "short response:") {
		return err1
	}
	return nil
}
