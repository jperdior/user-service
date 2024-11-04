package sns

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"

	"user-service/kit/event"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSBus struct {
	client   *sns.Client
	topicArn string
}

// NewSNSBus initializes a new SNSBus.
func NewSNSBus(topicArn, region string, endpoint *string) *SNSBus {

	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Error loading AWS config: %s\n", err)
	}

	awsConfig.Region = region

	if endpoint != nil && *endpoint != "" {
		awsConfig.BaseEndpoint = endpoint
	}

	client := sns.NewFromConfig(awsConfig)

	return &SNSBus{
		client:   client,
		topicArn: topicArn,
	}
}

// Publish publishes events to the SNS topic.
func (b *SNSBus) Publish(events []event.Event) error {
	for _, evt := range events {

		data, err := json.Marshal(evt.ToDTO())
		if err != nil {
			log.Printf("Error marshalling event data: %s - %s\n", evt.Type(), err)
			continue
		}

		envelope := event.EventEnvelope{
			EventType: evt.Type(),
			Data:      data,
		}
		message, err := json.Marshal(envelope)
		log.Printf("Marshalled event: %s\n", string(message))
		if err != nil {
			log.Printf("Error marshalling event: %s - %s\n", evt.Type(), err)
			continue
		}

		_, err = b.client.Publish(
			context.TODO(),
			&sns.PublishInput{
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
