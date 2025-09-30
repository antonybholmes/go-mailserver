package mailqueue

import (
	"sync"

	mailserver "github.com/antonybholmes/go-mailserver"
)

var (
	queue mailserver.MailQueue
	once  sync.Once
)

func InitMailQueue(q mailserver.MailQueue) {
	once.Do(func() {
		queue = q
	})
}

// SendMail adds an email to the mail queue
func SendMail(email *mailserver.MailItem) error {
	return queue.SendMail(email)
}
