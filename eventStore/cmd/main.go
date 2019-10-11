package main

import (
	"context"
	"eventStore/eventstore"
	"eventStore/kardex/usecase"
	"eventStore/messagingSystem"
	"log"
)

func main() {
	gRpcServer :=  eventstore.NewGRpcServer(":50051")
	natMessagingSystem := messagingSystem.NewNatsClient("nats://127.0.0.1:4222")
	useCaseKardex := usecase.NewKardexUseCase(2, natMessagingSystem, gRpcServer)
	if err := useCaseKardex.Publish(context.Background()); err != nil {
		log.Fatalf("error while start event store: %v",err)
	}
}
