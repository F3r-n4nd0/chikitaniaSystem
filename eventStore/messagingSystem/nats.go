package messagingSystem

import (
	"context"
	"eventStore/kardex"
	"eventStore/protobuf"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/nats-io/stan.go"
)

type natsClient struct {
	publisher message.Publisher
}

func NewNatsClient(url string) kardex.MessagingSystem {
	publisher, err := nats.NewStreamingPublisher(
		nats.StreamingPublisherConfig{
			ClusterID: "test-cluster",
			ClientID:  "kardex-store-api",
			StanOptions: []stan.Option{
				stan.NatsURL(url),
			},
			Marshaler: nats.GobMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}
	return &natsClient{
		publisher: publisher,
	}
}

func (n natsClient) Publish(ctx context.Context, event *protobuf.Event) error {
	msg := message.NewMessage(watermill.NewUUID(),[]byte(event.EventData) )
	if err := n.publisher.Publish("example.topic", msg); err != nil {
		return err
	}
	return nil
}
