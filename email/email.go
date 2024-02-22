package email

import (
	"net/mail"

	"github.com/antonybholmes/go-mailer"
	"github.com/antonybholmes/go-sys/env"
)

var emailer = mailer.NewSMTPMailer()

func init() {
	// force loading of enviromental variables if not done so
	env.Load()

	from := &mail.Address{Name: env.GetStr("NAME", ""), Address: env.GetStr("SMTP_FROM", "")}

	// Attempt to initialize by scanning enviromental variables.
	// If user has set them, magic, otherwise user will have to manually
	// specify
	emailer.
		SetUser(env.GetStr("SMTP_USER", ""), env.GetStr("SMTP_PASSWORD", "")).
		SetHost(env.GetStr("SMTP_HOST", ""), env.GetUint32("SMTP_PORT", 587)).
		SetFrom(from)
}

func SetUser(user string, password string) *mailer.SMTPMailer {
	return emailer.SetUser(user, password)
}

func SetHost(host string, port uint) *mailer.SMTPMailer {
	return emailer.SetHost(host, port)
}

func From() *mail.Address {
	return emailer.From()
}

func SetFrom(from *mail.Address) *mailer.SMTPMailer {
	return emailer.SetFrom(from)
}

func SendEmail(to *mail.Address, subject string, message string) error {
	return emailer.SendEmail(to, subject, message)
}

func SendHtmlEmail(to *mail.Address, subject string, message string) error {
	return emailer.SendHtmlEmail(to, subject, message)
}
