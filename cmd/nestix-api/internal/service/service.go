package service

import (
	"context"

	"github.com/edanko/nx/cmd/nestix-api/internal/adapters"
)

type Service struct {
	pathRepository         *adapters.PathRepository
	sheetPathRepository    *adapters.SheetPathRepository
	sheetPathDetRepository *adapters.SheetPathDetRepository
	orderRepository        *adapters.OrderRepository
	productRepository      *adapters.ProductRepository
	visualRepository       *adapters.VisualRepository
	machineRepository      *adapters.MachineRepository
}

func New(
	pathRepo *adapters.PathRepository,
	sheetPathRepo *adapters.SheetPathRepository,
	sheetPathDetRepo *adapters.SheetPathDetRepository,
	orderRepo *adapters.OrderRepository,
	productRepo *adapters.ProductRepository,
	visualRepo *adapters.VisualRepository,
	machineRepo *adapters.MachineRepository,
) *Service {
	return &Service{
		pathRepository:         pathRepo,
		sheetPathRepository:    sheetPathRepo,
		sheetPathDetRepository: sheetPathDetRepo,
		orderRepository:        orderRepo,
		productRepository:      productRepo,
		visualRepository:       visualRepo,
		machineRepository:      machineRepo,
	}
}

func (s Service) Something(ctx context.Context) {

}
