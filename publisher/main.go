package main

import (
	"context"
	// "encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/starRMS/pubsub-x-protobuff/pkg/ps"
	schema "github.com/starRMS/pubsub-x-protobuff/schema"
	"google.golang.org/protobuf/proto"
)

func main() {
	ps := ps.GetPubSub()

	t := ps.Client.Topic("topic")

	msg := &schema.SunTerraDeviceMessage{
		Name:     "Yes",
		Username: "Yes",
		Status:   1,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("unable to marshal proto")
	}

	result := t.Publish(context.TODO(), &pubsub.Message{
		Data:       data,
		Attributes: map[string]string{},
	})

	mID, err := result.Get(context.TODO())
	if err != nil {
		log.Fatal("failed to publish message")
	}

	log.Println("successfully published message with id", mID)
}
