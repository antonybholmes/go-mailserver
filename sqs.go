package mailserver

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/rs/zerolog/log"
)

var (
	AwsTypeString = aws.String("String")
	AwsTypeNumber = aws.String("Number")
	AwsTypeBinary = aws.String("Binary")
	AwsEmailValue = aws.String("email")
)

type SQSEmailQueue struct {
	client   *sqs.Client
	queueUrl *string
	ctx      context.Context
}

func NewSqsEmailQueue(queueUrl string) *SQSEmailQueue {
	ctx := context.Background()

	// Load AWS config (region and credentials)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal().Msgf("unable to load SDK config: %v", err)
	}

	client := sqs.NewFromConfig(cfg)

	return &SQSEmailQueue{client: client, queueUrl: aws.String(queueUrl), ctx: ctx}
}

func (queue *SQSEmailQueue) SendMail(mail *MailItem) error {

	body, err := json.Marshal(mail)

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
				DataType:    AwsTypeString,
				StringValue: AwsEmailValue,
			},
		},
	})

	return err
}
