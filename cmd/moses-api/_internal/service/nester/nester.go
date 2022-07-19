package nester

import (
	"context"

	"github.com/edanko/moses/internal/entities"
	"github.com/edanko/moses/internal/nester"
	"github.com/edanko/moses/internal/service/nest"
	"github.com/edanko/moses/internal/service/remnant"
	"github.com/edanko/moses/internal/service/spacing"
	"github.com/edanko/moses/internal/service/stock"
)

type UseCase interface {
	Nest(ctx context.Context, parts []*entities.Profile) ([]*entities.Nest, error)
	Renest(ctx context.Context, nests []*entities.Nest) ([]*entities.Nest, error)
}

type Service struct {
	nestService    nest.UseCase
	remnantService remnant.UseCase
	spacingService spacing.UseCase
	stockService   stock.UseCase
}

func NewService(n nest.UseCase, r remnant.UseCase, s spacing.UseCase, b stock.UseCase) *Service {
	return &Service{
		nestService:    n,
		remnantService: r,
		spacingService: s,
		stockService:   b,
	}
}

func (s *Service) Nest(ctx context.Context, machine entities.MachineType, parts []*entities.Profile) ([]*entities.Nest, error) {
	if machine == 0 {
		return nil, entities.ErrMachine
	}

	stockBar, err := s.stockService.GetOne(ctx, parts[0].Dim, parts[0].Quality)
	if err != nil {
		return nil, err
	}

	// if stockBar == nil {
	// 	 return nil, entities.ErrBars
	// }

	remnants, err := s.remnantService.GetNotUsed(ctx, parts[0].Project, parts[0].Dim, parts[0].Quality)
	if err != nil {
		return nil, err
	}

	bars := make([]*entities.Bar, 0, len(remnants))
	for _, r := range remnants {
		bars = append(bars, r.ToBar())
	}

	if len(parts) == 0 {
		return nil, entities.ErrParts
	}

	for _, part := range parts {
		l, err := s.spacingService.GetOne(ctx, machine, part.Dim, part.L)
		if err != nil {
			return nil, err
		}
		r, err := s.spacingService.GetOne(ctx, machine, part.Dim, part.R)
		if err != nil {
			return nil, err
		}

		part.Spacing = []float64{l.Length, r.Length}

		part.FullLength = l.Length + part.Length + r.Length
	}

	return nester.Nest(ctx, bars, stockBar.ToBar(), parts)
}

func (s *Service) Renest(ctx context.Context, machine entities.MachineType, nests []*entities.Nest) ([]*entities.Nest, error) {
	remnants := make([]*entities.Bar, 0, len(nests))
	parts := make([]*entities.Profile, 0)

	stockBar, err := s.stockService.GetOne(ctx, parts[0].Dim, parts[0].Quality)
	if err != nil {
		return nil, err
	}

	for _, n := range nests {
		remnants = append(remnants, n.Bar)
		parts = append(parts, n.Profiles...)

		err := s.nestService.Delete(ctx, n.ID)
		if err != nil {
			return nil, err
		}
	}

	for _, part := range parts {
		l, err := s.spacingService.GetOne(ctx, machine, part.Dim, part.L)
		if err != nil {
			return nil, err
		}
		r, err := s.spacingService.GetOne(ctx, machine, part.Dim, part.R)
		if err != nil {
			return nil, err
		}

		part.FullLength = l.Length + part.Length + r.Length
	}

	return nester.Nest(ctx, remnants, stockBar.ToBar(), parts)
}
