package usecase

import (
	"context"
	"errors"
	"eventStore/kardex"
	"eventStore/protobuf"
	"time"
)

type kardexUseCase struct {
	contextTimeout time.Duration
	messagingSystem kardex.MessagingSystem
	eventStoreServer kardex.EventStoreServer
}

func NewKardexUseCase(timeout time.Duration, messagingSystem kardex.MessagingSystem, eventStoreServer kardex.EventStoreServer) kardex.UseCase {
	return &kardexUseCase{
		contextTimeout: timeout,
		messagingSystem: messagingSystem,
		eventStoreServer: eventStoreServer,
	}
}

func (k kardexUseCase) Publish(ctx context.Context) error {
	if err := k.eventStoreServer.Listen(k); err != nil {
		return err
	}
	return nil
}

func (k kardexUseCase) GetEvents(context.Context, *protobuf.EventFilter) (*protobuf.EventResponse, error) {
	return nil, errors.New("operation GetEvents has not yet been implemented")
}

func (k kardexUseCase) CreateEvent(ctx context.Context, event *protobuf.Event) (*protobuf.Response, error) {
	if err := k.messagingSystem.Publish(ctx, event); err != nil {
		return nil , err
	}
	return &protobuf.Response{IsSuccess: true}, nil
}