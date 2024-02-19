package email

import (
	"github.com/antonybholmes/go-env"
	"github.com/antonybholmes/go-mailer"
)

var emailer = mailer.NewSMTPMailer()

func init() {
	// force loading of enviromental variables if not done so
	env.Load()

	from := mailer.NewMailbox(env.GetStr("NAME", ""), env.GetStr("SMTP_FROM", ""))

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

func From() *mailer.Mailbox {
	return emailer.From()
}

func SetFrom(from *mailer.Mailbox) *mailer.SMTPMailer {
	return emailer.SetFrom(from)
}

func SendEmail(to *mailer.Mailbox, subject string, message string) error {
	return emailer.SendEmail(to, subject, message)
}

func SendHtmlEmail(to *mailer.Mailbox, subject string, message string) error {
	return emailer.SendHtmlEmail(to, subject, message)
}
