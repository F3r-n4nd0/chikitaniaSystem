package kardex

import (
	"context"
	"kardexService/models"
)

type Repository interface {
	GetKardexById(ctx context.Context, kardexId string) (models.Kardex, error)
}