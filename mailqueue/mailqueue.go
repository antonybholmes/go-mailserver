package mailqueue

import (
	mailserver "github.com/antonybholmes/go-mailserver"
)

var queue mailserver.MailQueue

func Init(q mailserver.MailQueue) {
	queue = q
}

// SendMail adds an email to the mail queue
func SendMail(email *mailserver.MailItem) error {
	return queue.SendMail(email)
}
