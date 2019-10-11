package main

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/nats-io/stan.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kardexQueryStore/protobuf"
	"log"
	"time"
)

func main() {

	subscriber, err := nats.NewStreamingSubscriber(
		nats.StreamingSubscriberConfig{
			ClusterID:        "test-cluster",
			ClientID:         "example-subscriber",
			QueueGroup:       "example",
			DurableName:      "my-durable",
			SubscribersCount: 4, // how many goroutines should consume messages
			CloseTimeout:     time.Minute,
			AckWaitTimeout:   time.Second * 30,
			StanOptions: []stan.Option{
				stan.NatsURL("nats://127.0.0.1:4222"),
			},
			Unmarshaler: nats.GobMarshaler{},
		},
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		panic(err)
	}
	messages, err := subscriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}
	done := make(chan bool)
	go process(messages)
	<-done
}

func process(messages <-chan *message.Message) {
	for msg := range messages {

		kardexData := protobuf.RegisterKardexCommand{}
		err := json.Unmarshal(msg.Payload, &kardexData)
		if err != nil {
			panic(err)
		}
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			panic(err)
		}
		collection := client.Database("chikitania-db").Collection("kardex")
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.InsertOne(ctx, bson.M{"kardex_id": kardexData.KardexId, "product_id": kardexData.ProductId})
		log.Println("saved kardex: ", kardexData.KardexId, " id: ", res.InsertedID)
		if err != nil {
			log.Fatal(err)
		}
		msg.Ack()
	}
}
