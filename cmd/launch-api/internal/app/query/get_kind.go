package query

import (
	"context"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
	"github.com/google/uuid"
)

type GetKindRequest struct {
	ID uuid.UUID
}

type GetKindReadModel interface {
	Get(ctx context.Context, id uuid.UUID) (*kind.Kind, error)
}

type GetKindHandler struct {
	readModel GetKindReadModel
}

func (h GetKindHandler) Handle(ctx context.Context, query GetKindRequest) (tr KindModel, err error) {
	kind, err := h.readModel.Get(ctx, query.ID)
	if err != nil {
		return KindModel{}, err
	}
	ret := fromDomain(kind)

	return ret, nil
}

func NewGetKindHandler(
	readModel GetKindReadModel,
) GetKindHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetKindHandler{
		readModel: readModel,
	}
}
