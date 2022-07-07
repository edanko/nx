package machine

import (
	"context"

	"github.com/edanko/moses/internal/entities"
)

type Machine interface {
	Output(context.Context, *entities.Nest) ([]byte, error)
	Normalize(context.Context, *entities.Profile) (*entities.Profile, error)
}
