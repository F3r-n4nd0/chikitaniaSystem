package kardex


import (
"context"
"kardexService/models"
)

type UseCase interface {
	Register(ctx context.Context, newKardex models.NewKardex) error
	//GetById(ctx context.Context, kardexId string) (models.Kardex, error)
}

type QueryUseCase interface {
	GetById(ctx context.Context, kardexId string) (*models.Kardex, error)
}