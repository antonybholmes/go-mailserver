package mailer

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

const (
	QUEUE_EMAIL_CHANNEL = "email-channel"
)

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
