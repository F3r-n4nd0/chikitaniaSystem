package kardex

import (
	"context"
	"eventStore/protobuf"
)

type MessagingSystem interface {
	Publish(ctx context.Context, event *protobuf.Event) error
}

type EventStoreServer interface {
	Listen(srv protobuf.EventStoreServer) error
}