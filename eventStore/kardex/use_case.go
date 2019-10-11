package kardex

import (
	"context"
)

type UseCase interface {
	Publish(ctx context.Context) error
}
