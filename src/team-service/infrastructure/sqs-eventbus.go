package infrastructure

import (
	"errors"
	"fmt"
	"strings"
	"team-service/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// ErrQueueNotFound is returned when athe sent queue is not found.
var ErrQueueNotFound = errors.New("Specified queue not found")

// AmazonSqsEventBus is an event bus implementation using Amaazon SQS.
type AmazonSqsEventBus struct {
	svc       *sqs.SQS
	queueURLS []string
}

// NewAmazonSqsEventBus creates a instance of the AmazonSqsEventBus.
func NewAmazonSqsEventBus(requiredQueues []string) *AmazonSqsEventBus {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "league-manager-sqs"),
	}))

	svc := sqs.New(sess)

	queueUrls := make([]string, len(requiredQueues))

	if len(queueUrls) > 0 {
		for i, queueName := range requiredQueues {
			queueURLResponse, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
				QueueName: &queueName,
			})

			if err == nil {
				queueUrls[i] = *queueURLResponse.QueueUrl
			}
		}
	}

	return &AmazonSqsEventBus{
		svc:       svc,
		queueURLS: queueUrls,
	}
}

// Publish sends a new message to the event bus.
func (ev AmazonSqsEventBus) Publish(queueName string, evt domain.Event) error {
	requiredQueueURL := ""

	for _, queueURL := range ev.queueURLS {
		if strings.Contains(queueURL, queueName) {
			requiredQueueURL = queueURL
			break
		}
	}

	if len(requiredQueueURL) > 0 {

		result, err := ev.svc.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String(string(evt.AsEvent())),
			QueueUrl:    &requiredQueueURL,
		})

		if err != nil {
			fmt.Println("Error", err)
		}

		fmt.Println("Event published: ", *result.MessageId)

		return err
	}

	return ErrQueueNotFound
}
