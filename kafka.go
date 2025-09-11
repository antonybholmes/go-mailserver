package mailserver

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type KafkaEmailQueue struct {
	writer *kafka.Writer
}

func NewKafkaEmailQueue(writer *kafka.Writer) *KafkaEmailQueue {
	return &KafkaEmailQueue{writer: writer}
}

func (publisher *KafkaEmailQueue) SendEmail(email *QueueEmail) error {

	payload, err := json.Marshal(email)

	if err != nil {
		return err
	}

	// Sending a message to the topic
	message := kafka.Message{Key: []byte("email"),
		Value: payload,
	}

	// Write message to Kafka topic
	err = publisher.writer.WriteMessages(context.Background(), message)

	if err != nil {
		log.Debug().Msgf("could not write message: %v", err)
	}

	return nil
}
