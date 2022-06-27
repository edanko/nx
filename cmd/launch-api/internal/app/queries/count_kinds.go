package queries

import (
	"context"
)

type CountKindsRequest struct {
	Status *string
}

type CountKindsReadModel interface {
	Count(ctx context.Context, status *string) (int, error)
}

type CountKindsHandler struct {
	readModel CountKindsReadModel
}

func (h CountKindsHandler) Handle(ctx context.Context, query CountKindsRequest) (int, error) {
	totalKinds, err := h.readModel.Count(ctx, query.Status)
	if err != nil {
		return 0, err
	}

	return totalKinds, nil
}

func NewCountKindsHandler(
	readModel CountKindsReadModel,
) CountKindsHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return CountKindsHandler{
		readModel: readModel,
	}
}
