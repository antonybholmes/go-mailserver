package mailer

import (
	"bytes"
	"fmt"
	"net/mail"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

type SMTPMailer struct {
	from     *mail.Address
	user     string
	password string
	host     string
	addr     string
	port     uint
}

func NewSMTPMailer(user string, password string, host string, port uint, from *mail.Address) *SMTPMailer {
	//host := ""
	//port := uint(587)
	addr := fmt.Sprintf("%s:%d", host, port)

	log.Debug().Msgf("smtp: %s", addr)

	return &SMTPMailer{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		addr:     addr,
		from:     from}
}

// func (mailer *SMTPMailer) SetUser(user string, password string) *SMTPMailer {
// 	mailer.user = user
// 	mailer.password = password
// 	return mailer
// }

// func (mailer *SMTPMailer) SetHost(host string, port uint) *SMTPMailer {
// 	mailer.host = host
// 	mailer.port = port
// 	mailer.addr = fmt.Sprintf("%s:%d", mailer.host, mailer.port)
// 	return mailer
// }

// func (mailer *SMTPMailer) SetFrom(from *mail.Address) *SMTPMailer {
// 	mailer.from = from
// 	return mailer
// }

func (mailer *SMTPMailer) From() *mail.Address {
	return mailer.from
}

func (mailer *SMTPMailer) SendEmailRaw(to *mail.Address, body []byte) error {

	//from := os.Getenv("EMAIL")

	//password := os.Getenv("SMTP_PASSWORD")
	//user := os.Getenv("SMTP_USER")

	// Receiver email address.
	//to := "antony@antonyholmes.dev"

	// smtp server configuration.
	//smtpHost := os.Getenv("SMTP_HOST")
	//smtpPort := os.Getenv("SMTP_PORT")

	//addr := fmt.Sprintf("%s:%s", smtpHost, os.Getenv("SMTP_PORT"))

	//code := randomstring.CookieFriendlyString(32)

	// Message.
	// message := []byte(fmt.Sprintf("From: Experiment Database Service <%s>\r\n", mailer.from) +
	// 	fmt.Sprintf("To: %s\r\n", to) +
	// 	fmt.Sprintf("Subject: %s OTP code\r\n", os.Getenv("NAME")) +
	// 	"\r\n" +
	// 	fmt.Sprintf("Your one time code is: %s\r\n", code))

	// Authentication.
	auth := smtp.PlainAuth("", mailer.user, mailer.password, mailer.host)

	// Sending email.
	err := smtp.SendMail(mailer.addr, auth, mailer.from.Address, []string{
		to.Address,
	}, body)

	log.Debug().Msgf("aaaaa")

	if err != nil {
		log.Error().Msgf("%s", err)
		return err
	}

	log.Debug().Msgf("email to %s sent successfully", to.Address)

	return nil
}

func (mailer *SMTPMailer) SendEmail(to *mail.Address, subject string, message string) error {
	var body bytes.Buffer

	mailer.plainEmailHeader(to, subject, &body)
	emailLine(message, &body)

	return mailer.SendEmailRaw(to, body.Bytes())
}

func (mailer *SMTPMailer) SendHtmlEmail(to *mail.Address, subject string, message string) error {
	var body bytes.Buffer

	mailer.htmlEmailHeader(to, subject, &body)
	emailLine(message, &body)

	//log.Debug().Msgf("body %s", body.String())

	return mailer.SendEmailRaw(to, body.Bytes())
}

func (mailer *SMTPMailer) htmlEmailHeader(to *mail.Address, subject string, body *bytes.Buffer) {
	emailLine("MIME-version: 1.0", body)
	emailLine("Content-Type: text/html; charset=\"UTF-8\";", body)
	mailer.plainEmailHeader(to, subject, body)
}

func (mailer *SMTPMailer) plainEmailHeader(to *mail.Address, subject string, body *bytes.Buffer) {
	emailLine(fmt.Sprintf("From: %s", mailer.from), body)
	emailLine(fmt.Sprintf("To: %s", to), body)
	emailLine(fmt.Sprintf("Subject: %s", subject), body)
	// need empty line between header and body
	emailLine("", body)
}

func emailLine(s string, body *bytes.Buffer) {
	body.Write([]byte(fmt.Sprintf("%s\r\n", s)))
}
