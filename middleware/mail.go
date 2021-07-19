package middleware

import (
	"github.com/saicem/todo/global"
	"gopkg.in/gomail.v2"
)

func SendCaptchaEmail(toEmailAddress string, subject string, text string) error {
	config := global.Config.Email
	m := gomail.NewMessage()
	// 发件人
	m.SetHeader("From", config.FromEmail)
	// 收件人
	m.SetHeader("To", toEmailAddress)
	// 抄送
	// m.SetAddressHeader("Cc", "balabala@qq.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", "Hello, this is your captcha: "+text)
	// 附件
	// m.Attach()
	d := gomail.NewDialer(config.ServerHost, config.ServerPort, config.FromEmail, config.FromPassword)

	return d.DialAndSend(m)
}
