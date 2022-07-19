package nest

import (
	"context"

	"github.com/edanko/moses/internal/entities"
	"github.com/edanko/moses/internal/service/profile"
	"github.com/edanko/moses/internal/service/remnant"
	"github.com/edanko/moses/internal/service/spacing"

	"github.com/gofrs/uuid"
)

type Reader interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Nest, error)
	GetAll(ctx context.Context) ([]*entities.Nest, error)
	Get(ctx context.Context, project, dimension, quality string) ([]*entities.Nest, error)
}

type Writer interface {
	Create(ctx context.Context, e *entities.Nest) (*entities.Nest, error)
	Update(ctx context.Context, e *entities.Nest) (*entities.Nest, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type NestRepository interface {
	Reader
	Writer
}

type UseCase interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Nest, error)
	GetAll(ctx context.Context) ([]*entities.Nest, error)
	Get(ctx context.Context, project, dimension, quality string) ([]*entities.Nest, error)
	Create(ctx context.Context, e *entities.Nest) (*entities.Nest, error)
	Update(ctx context.Context, e *entities.Nest) (*entities.Nest, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type Service struct {
	repo           NestRepository
	remnantService remnant.UseCase
	profileService profile.UseCase
	spacingService spacing.UseCase
}

func NewService(n NestRepository, r remnant.UseCase, p profile.UseCase, s spacing.UseCase) *Service {
	return &Service{
		repo:           n,
		remnantService: r,
		profileService: p,
		spacingService: s,
	}
}

func (s *Service) Create(ctx context.Context, e *entities.Nest) (*entities.Nest, error) {
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

func (s *Service) GetOne(ctx context.Context, id uuid.UUID) (*entities.Nest, error) {
	n, err := s.repo.GetOne(ctx, id)
	if n == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (s *Service) Get(ctx context.Context, project, dimension, quality string) ([]*entities.Nest, error) {
	nests, err := s.repo.Get(ctx, project, dimension, quality)
	if err != nil {
		return nil, err
	}
	if len(nests) == 0 {
		return nil, entities.ErrNotFound
	}
	return nests, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*entities.Nest, error) {
	nests, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if len(nests) == 0 {
		return nil, entities.ErrNotFound
	}

	for _, n := range nests {
		for i, p := range n.ProfilesIds {
			profile, err := s.profileService.GetOne(ctx, p)
			if err != nil {
				return nil, err
			}

			profile.FullLength = n.Spacings[i*2] + profile.Length + n.Spacings[i*2+1]

			n.Profiles = append(n.Profiles, profile)
		}
		if n.Bar.RemnantID != nil {
			remnant, err := s.remnantService.GetOne(ctx, *n.Bar.RemnantID)
			if err != nil {
				return nil, err
			}
			n.Bar.Remnant = remnant
		}
	}

	return nests, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) DeleteAll(ctx context.Context) error {
	return s.repo.DeleteAll(ctx)
}

func (s *Service) Update(ctx context.Context, e *entities.Nest) (*entities.Nest, error) {
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	return s.repo.Update(ctx, e)
}
