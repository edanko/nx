package img

import (
	"context"

	"github.com/edanko/moses/internal/entities"
)

type UseCase interface {
	Output(context.Context, *entities.Nest) (string, error)
	Normalize(context.Context, *entities.Profile) (*entities.Profile, error)
}

type Machine struct {
}

func NewMachine() *Machine {
	return new(Machine)
}

func (m *Machine) Output(ctx context.Context, p *entities.Nest) (string, error) {
	return "", nil
}

func (m *Machine) Normalize(ctx context.Context, p *entities.Profile) (*entities.Profile, error) {
	return p, nil
}
