package infrastructure

import (
	"errors"
	"fmt"
	"strings"
	"team-service/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// ErrTopicNotFound is returned when the requested topic is not found.
var ErrTopicNotFound = errors.New("Specified topic not found")

// AmazonSnsEventBus is an event bus implementation using Amaazon SQS.
type AmazonSnsEventBus struct {
	svc       *sns.SNS
	availableTopics []string
}

// NewAmazonSnsEventBus creates a instance of the AmazonSnsEventBus.
func NewAmazonSnsEventBus() *AmazonSnsEventBus {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "league-manager-sqs"),
	}))

	svc := sns.New(sess)

	availableTopics, _ := svc.ListTopics(nil)

	availableTopicArns := make([]string, len(availableTopics.Topics))

	for i, t := range availableTopics.Topics {
		availableTopicArns[i] = *t.TopicArn
    }

	return &AmazonSnsEventBus{
		svc:       svc,
		availableTopics: availableTopicArns,
	}
}

// Publish sends a new message to the event bus.
func (ev AmazonSnsEventBus) Publish(publishTo string, evt domain.Event) error {
	requiredTopicArn := ""

	for _, t := range ev.availableTopics {
        if strings.Contains(t, publishTo) {
			requiredTopicArn = t
		}
    }

	if len(requiredTopicArn) > 0 {

		result, err := ev.svc.Publish(&sns.PublishInput{
			Message: aws.String(string(evt.AsEvent())),
			TopicArn:    aws.String(requiredTopicArn),
		})

		if err != nil {
			fmt.Println("Error", err)
		}

		fmt.Println("Event published: ", *result.MessageId)

		return err
	}

	return ErrTopicNotFound
}
