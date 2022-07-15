package queries

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
	"github.com/edanko/nx/pkg/pagination"
)

type ListKindsRequest struct {
	Limit  *int
	Cursor *string
	Status *string
}

type ListKindsReadModel interface {
	List(ctx context.Context, limit *int, createdAt *time.Time, id *uuid.UUID, status *string) ([]*kind.Kind, error)
}

type ListKindsHandler struct {
	readModel ListKindsReadModel
}

func (h ListKindsHandler) Handle(ctx context.Context, query ListKindsRequest) ([]KindModel, string, error) {
	var createdAt *time.Time
	var id *uuid.UUID
	if query.Cursor != nil {
		createdAtCursor, idCursor, err := pagination.DecodeCursor(*query.Cursor)
		if err != nil {
			return nil, "", errors.New("invalid-cursor")
		}
		createdAt = &createdAtCursor
		id = &idCursor
	}

	allKinds, err := h.readModel.List(ctx, query.Limit, createdAt, id, query.Status)
	if err != nil {
		return nil, "", err
	}

	ret := make([]KindModel, 0, len(allKinds))
	for _, k := range allKinds {
		ret = append(ret, fromDomain(k))
	}

	var nextCursor string
	if len(ret) > 0 {
		nextCursor = pagination.EncodeCursor(ret[len(ret)-1].CreatedAt, ret[len(ret)-1].ID)
	}

	return ret, nextCursor, nil
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
