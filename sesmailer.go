package mailer

import (
	"context"
	"net/mail"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/rs/zerolog/log"
)

type SesMailer struct {
	svc  *sesv2.Client
	from *string
}

func NewSesMailer(from *mail.Address) *SesMailer {
	// Create a new SES session
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal().Msgf("unable to load SDK config, %v", err)
	}

	// Create a new SES client
	svc := sesv2.NewFromConfig(cfg)

	return &SesMailer{
		svc:  svc,
		from: aws.String(from.String())}
}

// func (mailer *SMTPMailer) SetUser(user string, password string) *SMTPMailer {
// 	mailer.user = user
// 	mailer.password = password
// 	return mailer
// }

// func (mailer *SMTPMailer) SetHost(host string, port uint) *SMTPMailer {
// 	mailer.host = host
// 	mailer.port = port
// 	mailer.addr = fmt.Sprintf("%s:%d", mailer.host, mailer.port)
// 	return mailer
// }

// func (mailer *SMTPMailer) SetFrom(from *mail.Address) *SMTPMailer {
// 	mailer.from = from
// 	return mailer
// }

func (mailer *SesMailer) SendHtmlEmail(to *mail.Address, subject string, message string, html string) error {
	input := &sesv2.SendEmailInput{
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Html: &types.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(html),
					},
					// Text: &types.Content{
					// 	Charset: aws.String("UTF-8"),
					// 	Data:    aws.String("Hello, world!\nThis is a test email sent with AWS SES V2."),
					// },
				},
				Subject: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(subject),
				},
			},
		},
		Destination: &types.Destination{
			ToAddresses: []string{to.Address}, // Replace with the recipient's email address
		},
		FromEmailAddress: mailer.from, // Replace with your verified sender email address
	}

	// Send the email
	resp, err := mailer.svc.SendEmail(context.TODO(), input)
	if err != nil {
		return err
	}

	// Print the message ID on success
	log.Debug().Msgf("Email sent! Message ID: %s", *resp.MessageId)

	return nil
}
