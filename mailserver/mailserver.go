package mailserver

import (
	"net/mail"

	"github.com/antonybholmes/go-mailer"
	"github.com/antonybholmes/go-sys/env"
	"github.com/rs/zerolog/log"
)

var instance *mailer.SMTPMailer

func Init() {
	// force loading of enviromental variables if not done so
	//env.Reload()

	from := &mail.Address{Name: env.GetStr("NAME", ""), Address: env.GetStr("SMTP_FROM", "")}

	log.Debug().Msgf("port %s", env.GetStr("SMTP_HOST", ""))
	// Attempt to initialize by scanning enviromental variables.
	// If user has set them, magic, otherwise user will have to manually
	// specify
	// instance.
	// 	SetUser(env.GetStr("SMTP_USER", ""), env.GetStr("SMTP_PASSWORD", "")).
	// 	SetHost(env.GetStr("SMTP_HOST", ""), env.GetUint32("SMTP_PORT", 587)).
	// 	SetFrom(from)

	instance = mailer.NewSMTPMailer(env.GetStr("SMTP_USER", ""),
		env.GetStr("SMTP_PASSWORD", ""),
		env.GetStr("SMTP_HOST", ""),
		env.GetUint32("SMTP_PORT", 587),
		from)
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

func From() *mail.Address {
	return instance.From()
}

func SendEmail(to *mail.Address, subject string, message string) error {
	return instance.SendEmail(to, subject, message)
}

func SendHtmlEmail(to *mail.Address, subject string, message string) error {
	return instance.SendHtmlEmail(to, subject, message)
}
