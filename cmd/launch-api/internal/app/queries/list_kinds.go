package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
)

type ListKindsRequest struct {
	Limit  *int
	After  *uuid.UUID
	Status *string
}

type ListKindsReadModel interface {
	List(ctx context.Context, limit *int, after *uuid.UUID, status *string) ([]*kind.Kind, error)
}

type ListKindsHandler struct {
	readModel ListKindsReadModel
}

func (h ListKindsHandler) Handle(ctx context.Context, query ListKindsRequest) ([]KindModel, error) {
	allKinds, err := h.readModel.List(ctx, query.Limit, query.After, query.Status)
	if err != nil {
		return nil, err
	}

	ret := make([]KindModel, 0, len(allKinds))
	for _, k := range allKinds {
		ret = append(ret, fromDomain(k))
	}

	return ret, nil
}

func NewListKindsHandler(
	readModel ListKindsReadModel,
) ListKindsHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return ListKindsHandler{
		readModel: readModel,
	}
}
