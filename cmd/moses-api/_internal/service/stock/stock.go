package stock

import (
	"context"
	"fmt"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
)

type Reader interface {
	GetOne(ctx context.Context, dim, quality string) (*entities.Stock, error)
	GetAll(ctx context.Context) ([]*entities.Stock, error)
}

type Writer interface {
	Create(ctx context.Context, e *entities.Stock) (*entities.Stock, error)
	Update(ctx context.Context, e *entities.Stock) (*entities.Stock, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type StockRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetOne(ctx context.Context, dim, quality string) (*entities.Stock, error)
	GetAll(ctx context.Context) ([]*entities.Stock, error)
	Create(ctx context.Context, e *entities.Stock) (*entities.Stock, error)
	Update(ctx context.Context, e *entities.Stock) (*entities.Stock, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type Service struct {
	repo StockRepository
}

func NewService(r StockRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, e *entities.Stock) (*entities.Stock, error) {
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

func (s *Service) GetOne(ctx context.Context, dim, quality string) (*entities.Stock, error) {
	r, err := s.repo.GetOne(ctx, dim, quality)
	if r == nil {
		return nil, fmt.Errorf("stock bar length for dim \"%s\" and quality \"%s\"", dim, quality)
	}
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*entities.Stock, error) {
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

func (s *Service) Update(ctx context.Context, e *entities.Stock) (*entities.Stock, error) {
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, e)
}
