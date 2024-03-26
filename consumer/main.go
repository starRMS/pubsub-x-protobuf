package main

import (
	"context"
	"encoding/binary"

	// "encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/starRMS/pubsub-x-protobuff/pkg/ps"
	schema "github.com/starRMS/pubsub-x-protobuff/schema"
	"google.golang.org/protobuf/proto"
)

func main() {
	ps := ps.GetPubSub()

	s := ps.Client.Subscription("topic")

	s.ReceiveSettings.Synchronous = true
	s.ReceiveSettings.NumGoroutines = 1
	s.Receive(context.TODO(), func(ctx context.Context, m *pubsub.Message) {
		var msg schema.SunTerraDeviceMessage

		log.Println("Received size", len(m.Data))
		log.Println("Received size", binary.Size(m.Data))

		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			log.Println(err)
			m.Nack()
		} else {
			log.Printf("Berhasil; Name: %s, Username: %s\n", msg.Name, msg.Username)
			m.Ack()
		}
	})
}
