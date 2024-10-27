package sns

import (
	"encoding/json"
	"log"

	"user-service/kit/event"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSBus struct {
	client   *sns.SNS
	topicArn string
}

// NewSNSBus initializes a new SNSBus.
func NewSNSBus(topicArn, region string, endpoint *string) *SNSBus {
	awsConfig := &aws.Config{
		Region: aws.String(region),
	}

	if endpoint != nil && *endpoint != "" {
		awsConfig.Endpoint = aws.String(*endpoint)
	}

	sess := session.Must(session.NewSession(awsConfig))
	client := sns.New(sess)

	return &SNSBus{
		client:   client,
		topicArn: topicArn,
	}
}

// Publish publishes events to the SNS topic.
func (b *SNSBus) Publish(events []event.Event) error {
	for _, evt := range events {
		message, err := json.Marshal(evt)
		if err != nil {
			log.Printf("Error marshalling event: %s - %s\n", evt.Type(), err)
			continue
		}

		_, err = b.client.Publish(&sns.PublishInput{
			Message:  aws.String(string(message)),
			TopicArn: aws.String(b.topicArn),
		})
		if err != nil {
			log.Printf("Error publishing to SNS: %s - %s\n", evt.Type(), err)
		} else {
			log.Printf("Published event %s to SNS\n", evt.Type())
		}
	}
	return nil
}
