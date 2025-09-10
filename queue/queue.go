package queue

import (
	mailer "github.com/antonybholmes/go-mailer"
)

var queue mailer.EmailQueue

func Init(q mailer.EmailQueue) {
	queue = q
}

func PublishEmail(email *mailer.QueueEmail) error {
	return queue.SendEmail(email)
}
