package spacing

import (
	"context"
	"errors"
	"fmt"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
)

type Reader interface {
	GetAll(ctx context.Context, machine entities.MachineType) ([]*entities.Spacing, error)
	GetOne(ctx context.Context, machine entities.MachineType, dim string, e *entities.End) (*entities.Spacing, error)
}

type Writer interface {
	Create(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error)
	Update(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type SpacingRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAll(ctx context.Context, machine entities.MachineType) ([]*entities.Spacing, error)
	GetOne(ctx context.Context, machine entities.MachineType, dim string, e *entities.End) (*entities.Spacing, error)
	Create(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error)
	Update(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type Service struct {
	repo SpacingRepository
}

func NewService(r SpacingRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error) {
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	n, err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (s *Service) GetOne(ctx context.Context, machine entities.MachineType, dim string, e *entities.End) (*entities.Spacing, error) {
	sp, err := s.repo.GetOne(ctx, machine, dim, e)
	if err != nil && errors.Is(err, entities.ErrNotFound) {
		bevel := e.WebBevel != nil || e.FlangeBevel != nil
		scallop := e.Scallop != nil
		return nil, fmt.Errorf("spacing for machine \"%s\", dim \"%s\", name \"%s\", bevel \"%t\", scallop \"%t\" not found", machine, dim, e.Name, bevel, scallop)
	} else if err != nil {
		return nil, err
	}
	return sp, nil
}

func (s *Service) GetAll(ctx context.Context, machine entities.MachineType) ([]*entities.Spacing, error) {
	spacings, err := s.repo.GetAll(ctx, machine)
	if err != nil {
		return nil, err
	}
	if len(spacings) == 0 {
		return nil, entities.ErrNotFound
	}
	return spacings, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) DeleteAll(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}

func (s *Service) Update(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error) {
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, e)
}
