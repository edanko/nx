package query

import (
	"context"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
)

type GetKindByNameRequest struct {
	Name string
}

type GetKindByNameReadModel interface {
	GetByName(ctx context.Context, name string) (*kind.Kind, error)
}

type GetKindByNameHandler struct {
	readModel GetKindByNameReadModel
}

func (h GetKindByNameHandler) Handle(ctx context.Context, query GetKindByNameRequest) (tr KindModel, err error) {
	kind, err := h.readModel.GetByName(ctx, query.Name)
	if err != nil {
		return KindModel{}, err
	}
	ret := fromDomain(kind)

	return ret, nil
}

func NewGetKindByNameHandler(
	readModel GetKindByNameReadModel,
) GetKindByNameHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return GetKindByNameHandler{
		readModel: readModel,
	}
}
