package usecase

import (
	"context"
	"github.com/google/uuid"
	"kardexService/kardex"
	"kardexService/models"
	"kardexService/protobuf"
	"time"
)

type kardexUseCase struct {
	contextTimeout   time.Duration
	eventStoreClient kardex.EventStoreClient
}

func NewKardexUseCase(timeout time.Duration, eventStoreClient kardex.EventStoreClient) kardex.UseCase {
	return &kardexUseCase{
		contextTimeout: timeout,
		eventStoreClient: eventStoreClient,
	}
}

func (k kardexUseCase) Register(ctx context.Context, newKardex models.NewKardex) error {
	var registerKardexCommand protobuf.RegisterKardexCommand
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	// ADD product ID validation
	registerKardexCommand.KardexId = newUUID.String()
	registerKardexCommand.ProductId = newKardex.ProductId
	registerKardexCommand.CreatedOn = time.Now().Unix()
	if err := k.eventStoreClient.Register(ctx, registerKardexCommand); err != nil {
		return err
	}
	return nil
}
