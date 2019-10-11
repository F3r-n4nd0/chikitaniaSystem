package usecase

import (
	"context"
	"kardexService/kardex"
	"kardexService/models"
	"time"
)

type kardexQueryUseCase struct {
	contextTimeout   time.Duration
	repository kardex.Repository
}

func NewQueryKardexUseCase(timeout time.Duration, repository kardex.Repository) kardex.QueryUseCase {
	return &kardexQueryUseCase{
		contextTimeout: timeout,
		repository: repository,
	}
}

func (k kardexQueryUseCase) GetById(ctx context.Context, kardexId string) (*models.Kardex, error) {
	kardex, err := k.repository.GetKardexById(ctx, kardexId)
	if err != nil {
		return nil, err
	}
	return &kardex, nil
}