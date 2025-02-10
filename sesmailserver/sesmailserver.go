package sesmailserver

import (
	"net/mail"

	"github.com/antonybholmes/go-mailer"
)

var instance *mailer.SesMailer

func Init(from *mail.Address) {
	//from := &mail.Address{Name: env.GetStr("NAME", ""), Address: env.GetStr("SMTP_FROM", "")}

	instance = mailer.NewSesMailer(from)
}

// func SetUser(user string, password string) *mailer.SMTPMailer {
// 	return instance.SetUser(user, password)
// }

// func SetHost(host string, port uint) *mailer.SMTPMailer {
// 	return instance.SetHost(host, port)
// }

// func SetFrom(from *mail.Address) *mailer.SMTPMailer {
// 	return instance.SetFrom(from)
// }

// func SendEmail(to *mail.Address, subject string, message string) error {
// 	return instance.SendEmail(to, subject, message)
// }

func SendHtmlEmail(to *mail.Address, subject string, message string) error {
	return instance.SendHtmlEmail(to, subject, "", message)
}
