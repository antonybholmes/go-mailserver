package queue

import (
	mailserver "github.com/antonybholmes/go-mail-server"
)

var queue mailserver.EmailQueue

func Init(q mailserver.EmailQueue) {
	queue = q
}

func PublishEmail(email *mailserver.QueueEmail) error {
	return queue.SendEmail(email)
}
