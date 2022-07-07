package remnant

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
)

type Reader interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Remnant, error)
	GetAll(ctx context.Context) ([]*entities.Remnant, error)
	GetNotUsed(ctx context.Context, project, dimension, quality string) ([]*entities.Remnant, error)
	GetAllNotUsed(ctx context.Context) ([]*entities.Remnant, error)
}

type Writer interface {
	Create(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error)
	Update(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type RemnantRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Remnant, error)
	GetAll(ctx context.Context) ([]*entities.Remnant, error)
	GetNotUsed(ctx context.Context, project, dimension, quality string) ([]*entities.Remnant, error)
	GetAllNotUsed(ctx context.Context) ([]*entities.Remnant, error)
	Create(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error)
	Update(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type Service struct {
	repo RemnantRepository
}

func NewService(r RemnantRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error) {
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	r, err := s.repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Service) GetOne(ctx context.Context, id uuid.UUID) (*entities.Remnant, error) {
	r, err := s.repo.GetOne(ctx, id)
	if r == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Service) GetNotUsed(ctx context.Context, project, dimension, quality string) ([]*entities.Remnant, error) {
	remnants, err := s.repo.GetNotUsed(ctx, project, dimension, quality)
	if err != nil {
		return nil, err
	}
	// if len(remnants) == 0 {
	// 	 return nil, entities.ErrNotFound
	// }
	return remnants, nil
}

func (s *Service) GetAllNotUsed(ctx context.Context) ([]*entities.Remnant, error) {
	remnants, err := s.repo.GetAllNotUsed(ctx)
	if err != nil {
		return nil, err
	}
	if len(remnants) == 0 {
		return nil, entities.ErrNotFound
	}
	return remnants, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*entities.Remnant, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, entities.ErrNotFound
	}
	return users, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) DeleteAll(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}

func (s *Service) Update(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error) {
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, e)
}
