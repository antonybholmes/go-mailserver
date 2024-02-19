package email

import (
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

type Mailbox struct {
	name  string
	email string
}

func NewMailbox(name string, email string) *Mailbox {
	return &Mailbox{name, email}
}

func (mailbox *Mailbox) String() string {
	return fmt.Sprintf("%s <%s>", mailbox.name, mailbox.email)
}

type SMTPMailer struct {
	user     string
	password string
	host     string
	port     uint
	addr     string
	from     *Mailbox
}

func NewSMTPMailer() *SMTPMailer {
	host := ""
	port := uint(587)
	addr := fmt.Sprintf("%s:%d", host, port)

	return &SMTPMailer{
		user:     "",
		password: "",
		host:     host,
		port:     port,
		addr:     addr,
		from:     nil}
}

func (mailer *SMTPMailer) SetUser(user string, password string) *SMTPMailer {
	mailer.user = user
	mailer.password = password
	return mailer
}

func (mailer *SMTPMailer) SetHost(host string, port uint) *SMTPMailer {
	mailer.host = host
	mailer.port = port
	mailer.addr = fmt.Sprintf("%s:%d", mailer.host, mailer.port)
	return mailer
}

func (mailer *SMTPMailer) From() *Mailbox {
	return mailer.from
}

func (mailer *SMTPMailer) SetFrom(from *Mailbox) *SMTPMailer {
	mailer.from = from
	return mailer
}

func (mailer *SMTPMailer) SendEmailRaw(to *Mailbox, body []byte) error {

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
	err := smtp.SendMail(mailer.addr, auth, mailer.from.email, []string{
		to.email,
	}, body)

	if err != nil {
		log.Error().Msgf("%s", err)
		return err
	}

	log.Debug().Msgf("Email Sent Successfully!")

	return nil
}

func (mailer *SMTPMailer) SendEmail(to *Mailbox, subject string, message string) error {
	var body bytes.Buffer

	mailer.plainEmailHeader(to, subject, &body)
	EmailLine(message, &body)

	return mailer.SendEmailRaw(to, body.Bytes())
}

func (mailer *SMTPMailer) SendHtmlEmail(to *Mailbox, subject string, message string) error {
	var body bytes.Buffer

	mailer.htmlEmailHeader(to, subject, &body)
	EmailLine(message, &body)

	return mailer.SendEmailRaw(to, body.Bytes())
}

func (mailer *SMTPMailer) htmlEmailHeader(to *Mailbox, subject string, body *bytes.Buffer) {
	EmailLine("MIME-version: 1.0", body)
	EmailLine("Content-Type: text/plain; charset=\"UTF-8\";", body)
	mailer.plainEmailHeader(to, subject, body)
}

func (mailer *SMTPMailer) plainEmailHeader(to *Mailbox, subject string, body *bytes.Buffer) {
	FromHeader(mailer.from, body)
	ToHeader(to, body)
	SubjectHeader(subject, body)
	// need empty line between header and body
	EmailLine("", body)
}

func FromHeader(mailbox *Mailbox, body *bytes.Buffer) {
	EmailLine(fmt.Sprintf("From: %s", mailbox), body)
}

func ToHeader(mailbox *Mailbox, body *bytes.Buffer) {
	EmailLine(fmt.Sprintf("To: %s", mailbox), body)
}

func SubjectHeader(subject string, body *bytes.Buffer) {
	EmailLine(fmt.Sprintf("Subject: %s", subject), body)
}

func EmailLine(s string, body *bytes.Buffer) {
	body.Write([]byte(fmt.Sprintf("%s\r\n", s)))
}
