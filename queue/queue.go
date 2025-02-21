package queue

import (
	mailer "github.com/antonybholmes/go-mailer"
)

var publisher mailer.EmailPublisher

func Init(p mailer.EmailPublisher) {
	publisher = p
}

func PublishEmail(email *mailer.QueueEmail) error {
	return publisher.PublishEmail(email)
}
