package mailer

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const (
	QUEUE_EMAIL_CHANNEL = "email-channel"
)

const (
	QUEUE_EMAIL_TYPE_VERIFY           = "verify"
	QUEUE_EMAIL_TYPE_VERIFIED         = "verified"
	QUEUE_EMAIL_TYPE_PASSWORDLESS     = "passwordless"
	QUEUE_EMAIL_TYPE_PASSWORD_RESET   = "password-reset"
	QUEUE_EMAIL_TYPE_PASSWORD_UPDATED = "password-updated"
	QUEUE_EMAIL_TYPE_EMAIL_RESET      = "email-reset"
	QUEUE_EMAIL_TYPE_EMAIL_UPDATED    = "email-updated"
	QUEUE_EMAIL_TYPE_ACCOUNT_CREATED  = "account-created"
	QUEUE_EMAIL_TYPE_ACCOUNT_UPDATED  = "account-updated"
)

type QueueEmail struct {
	Name      string `json:"name"`
	To        string `json:"to"`
	EmailType string `json:"type"`
	Ttl       string `json:"ttl"`
	Token     string `json:"token"`
	//RedirectUrl string `json:"redirectUrl"`
	LinkUrl string `json:"linkUrl"`
	// Used to indicate if emails should be for web users or
	// geared towards API users who just want the tokens without
	// any links
	Mode string `json:"mode"`
}

type EmailPublisher interface {
	PublishEmail(email *QueueEmail) error
}

type RedisEmailPublisher struct {
	rdb *redis.Client
}

func NewRedisEmailPublisher(rdb *redis.Client) *RedisEmailPublisher {
	return &RedisEmailPublisher{rdb: rdb}
}

func (publisher *RedisEmailPublisher) PublishEmail(email *QueueEmail) error {
	payload, err := json.Marshal(email)

	if err != nil {
		return err
	}

	return publisher.publish(QUEUE_EMAIL_CHANNEL, payload)
}

func (publisher *RedisEmailPublisher) publish(channel string, data []byte) error {
	//log.Debug().Msgf("send %v", data)
	return publisher.rdb.Publish(context.Background(), channel, data).Err()
}

type KafkaEmailPublisher struct {
	writer *kafka.Writer
}

func NewKafkaEmailPublisher(writer *kafka.Writer) *KafkaEmailPublisher {
	return &KafkaEmailPublisher{writer: writer}
}

func (publisher *KafkaEmailPublisher) PublishEmail(email *QueueEmail) error {

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
