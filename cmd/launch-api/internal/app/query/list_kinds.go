package query

import (
	"context"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
)

type ListKindsRequest struct {
	Limit  *int
	Offset *int
	Status *string
}

type ListKindsReadModel interface {
	List(ctx context.Context, limit *int, offset *int, status *string) ([]*kind.Kind, error)
}

type ListKindsHandler struct {
	readModel ListKindsReadModel
}

func (h ListKindsHandler) Handle(ctx context.Context, query ListKindsRequest) ([]KindModel, error) {
	allKinds, err := h.readModel.List(ctx, query.Limit, query.Offset, query.Status)
	if err != nil {
		return nil, err
	}

	ret := make([]KindModel, 0, len(allKinds))
	for _, kind := range allKinds {
		ret = append(ret, fromDomain(kind))
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
