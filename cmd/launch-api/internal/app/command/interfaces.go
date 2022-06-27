package command

import (
	"context"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
	"github.com/google/uuid"
)

type eventBus interface {
	Publish(ctx context.Context, event any) error
}

type sanitazer interface {
	Sanitize(s string) string
}

type KindRepository interface {
	Create(ctx context.Context, k *kind.Kind) error
	Get(ctx context.Context, id uuid.UUID) (*kind.Kind, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Exist(ctx context.Context, name string) (bool, error)
	Update(ctx context.Context, k *kind.Kind) error

	List(ctx context.Context, limit *int, offset *int, status *string) ([]*kind.Kind, error)
	Count(ctx context.Context, status *string) (int, error)
	GetByName(ctx context.Context, name string) (*kind.Kind, error)
}
