package mailer

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const (
	QUEUE_EMAIL_CHANNEL = "email-channel"
)

var (
	AWSStringType = aws.String("String")
	AWSNumberType = aws.String("Number")
	AWSBinaryType = aws.String("Binary")
	AWSEmailValue = aws.String("email")
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
	QUEUE_EMAIL_TYPE_OTP              = "otp"
)

type QueueEmail struct {
	Name      string `json:"name"`
	To        string `json:"to"`
	EmailType string `json:"type"`
	TTL       string `json:"ttl"`
	Token     string `json:"token"`
	//RedirectUrl string `json:"redirectUrl"`
	LinkUrl string `json:"linkUrl"`
	// Used to indicate if emails should be for web users or
	// geared towards API users who just want the tokens without
	// any links
	Mode string `json:"mode"`
}

type EmailQueue interface {
	SendEmail(email *QueueEmail) error
}

type RedisEmailQueue struct {
	rdb *redis.Client
}

func NewRedisEmailQueue(rdb *redis.Client) *RedisEmailQueue {
	return &RedisEmailQueue{rdb: rdb}
}

func (publisher *RedisEmailQueue) SendEmail(email *QueueEmail) error {
	payload, err := json.Marshal(email)

	if err != nil {
		return err
	}

	return publisher.publish(QUEUE_EMAIL_CHANNEL, payload)
}

func (publisher *RedisEmailQueue) publish(channel string, data []byte) error {
	//log.Debug().Msgf("send %v", data)
	return publisher.rdb.Publish(context.Background(), channel, data).Err()
}

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

type SQSEmailQueue struct {
	client   *sqs.Client
	queueUrl *string
	ctx      context.Context
}

func NewSQSEmailQueue(queueUrl string) *SQSEmailQueue {
	ctx := context.Background()

	// Load AWS config (region and credentials)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal().Msgf("unable to load SDK config: %v", err)
	}

	client := sqs.NewFromConfig(cfg)

	return &SQSEmailQueue{client: client, queueUrl: aws.String(queueUrl), ctx: ctx}
}

func (queue *SQSEmailQueue) SendEmail(email *QueueEmail) error {

	body, err := json.Marshal(email)

	if err != nil {
		return err
	}

	// Send to SQS
	_, err = queue.client.SendMessage(queue.ctx, &sqs.SendMessageInput{
		QueueUrl:    queue.queueUrl,
		MessageBody: aws.String(string(body)),
		// Optional: add message attributes
		MessageAttributes: map[string]types.MessageAttributeValue{
			"type": {
				DataType:    AWSStringType,
				StringValue: AWSEmailValue,
			},
		},
	})

	return err
}
