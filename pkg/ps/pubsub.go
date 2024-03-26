package ps

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type PubSub struct {
	projectID string
	Client    *pubsub.Client
}

func GetPubSub() *PubSub {
	ps := &PubSub{}

	// Set emulator pubsub
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	ps.projectID = "project-id"

	client, err := pubsub.NewClient(context.Background(), ps.projectID, option.WithoutAuthentication())
	if err != nil {
		log.Fatal("unable to create pubsub client", err)
	}

	ps.Client = client

	ps.CreateSubscription(ps.CreateTopic())

	return ps
}

func (ps *PubSub) CreateTopic() (topic *pubsub.Topic) {
	var err error
	topic, err = ps.Client.CreateTopicWithConfig(context.TODO(), "topic", &pubsub.TopicConfig{
		// RetentionDuration: 15 * 24 * time.Hour,
	})
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatal("unable to create topic", err)
		} else {
			log.Println("topic already exists")
			topic = ps.Client.Topic("topic")
		}
	}

	return
}

func (ps *PubSub) CreateSubscription(topic *pubsub.Topic) (sub *pubsub.Subscription) {
	var err error
	sub, err = ps.Client.CreateSubscription(context.TODO(), "topic", pubsub.SubscriptionConfig{
		Topic:                     topic,
		EnableExactlyOnceDelivery: true,
	})
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatal("unable to create sub", err)
		} else {
			log.Println("subscription already exists")
			sub = ps.Client.Subscription("topic")
		}
	}

	return
}
