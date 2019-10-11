package kardex

import (
	"context"
	"kardexService/protobuf"
)

type EventStoreClient interface {
	Register(ctx context.Context, newKardex protobuf.RegisterKardexCommand) error
}