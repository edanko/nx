package profile

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
)

type Reader interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Profile, error)
	GetAll(ctx context.Context) ([]*entities.Profile, error)
	GetAllLaunch(ctx context.Context, launch string) ([]*entities.Profile, error)
	Get(ctx context.Context, project, dimension, quality string) ([]*entities.Profile, error)
}

type Writer interface {
	Create(ctx context.Context, e *entities.Profile) (*entities.Profile, error)
	Update(ctx context.Context, e *entities.Profile) (*entities.Profile, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type ProfileRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Profile, error)
	GetAll(ctx context.Context) ([]*entities.Profile, error)
	GetAllLaunch(ctx context.Context, launch string) ([]*entities.Profile, error)
	Get(ctx context.Context, project, dimension, quality string) ([]*entities.Profile, error) // get free?
	Create(ctx context.Context, e *entities.Profile) (*entities.Profile, error)
	Update(ctx context.Context, e *entities.Profile) (*entities.Profile, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type Service struct {
	repo ProfileRepository
}

func NewService(r ProfileRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, e *entities.Profile) (*entities.Profile, error) {
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

func (s *Service) GetOne(ctx context.Context, id uuid.UUID) (*entities.Profile, error) {
	r, err := s.repo.GetOne(ctx, id)
	if r == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Service) Get(ctx context.Context, project, dimension, quality string) ([]*entities.Profile, error) {
	profiles, err := s.repo.Get(ctx, project, dimension, quality)
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, entities.ErrNotFound
	}
	return profiles, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*entities.Profile, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, entities.ErrNotFound
	}
	return users, nil
}

func (s *Service) GetAllLaunch(ctx context.Context, launch string) ([]*entities.Profile, error) {
	users, err := s.repo.GetAllLaunch(ctx, launch)
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

func (s *Service) Update(ctx context.Context, e *entities.Profile) (*entities.Profile, error) {
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, e)
}
