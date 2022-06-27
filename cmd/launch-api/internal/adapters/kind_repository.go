package adapters

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/nx/cmd/launch-api/internal/adapters/repositories/ent"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/repositories/ent/kind"
	domain "github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
)

type KindRepository struct {
	client *ent.KindClient
}

func NewKindRepository(c *ent.KindClient) *KindRepository {
	return &KindRepository{
		client: c,
	}
}

func (r *KindRepository) Create(
	ctx context.Context,
	k *domain.Kind,
) error {
	return r.client.
		Create().
		SetID(k.ID()).
		SetName(k.Name()).
		SetNillableDescription(k.Description()).
		SetStatus(kind.Status(k.Status().String())).
		Exec(ctx)
}

func (r *KindRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Kind, error) {
	e, err := r.client.
		Query().
		Where(kind.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapEntity(e), nil
}

func (r *KindRepository) GetByName(
	ctx context.Context,
	name string,
) (*domain.Kind, error) {
	e, err := r.client.
		Query().
		Where(kind.NameEQ(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapEntity(e), nil
}

func (r *KindRepository) List(
	ctx context.Context,
	limit *int,
	after *uuid.UUID,
	status *string,
) ([]*domain.Kind, error) {
	kindQuery := r.client.Query().
		Order(ent.Asc(kind.FieldName))

	if limit != nil {
		kindQuery = kindQuery.Limit(*limit)
	}
	if after != nil {
		kindQuery = kindQuery.Where(kind.IDGT(*after))
	}
	if status != nil {
		kindQuery = kindQuery.Where(
			kind.StatusEQ(kind.Status(*status)),
		)
	}

	es, err := kindQuery.All(ctx)
	if err != nil {
		return nil, err
	}

	return mapEntities(es), nil
}

func (r *KindRepository) Count(
	ctx context.Context,
	status *string,
) (int, error) {
	kindQuery := r.client.Query()
	if status != nil {
		kindQuery = kindQuery.Where(
			kind.StatusEQ(kind.Status(*status)),
		)
	}
	return kindQuery.Count(ctx)
}

func (r *KindRepository) Update(
	ctx context.Context,
	k *domain.Kind,
) error {
	return r.client.
		UpdateOneID(k.ID()).
		SetName(k.Name()).
		SetNillableDescription(k.Description()).
		SetStatus(kind.Status(k.Status().String())).
		Exec(ctx)
}

func (r *KindRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.client.DeleteOneID(id).Exec(ctx)
}

func (r *KindRepository) Exist(
	ctx context.Context,
	name string,
) (bool, error) {
	return r.client.Query().
		Where(kind.NameEQ(name)).
		Exist(ctx)
}

func mapEntity(e *ent.Kind) *domain.Kind {
	kind, _ := domain.NewKind(
		e.ID,
		e.Name,
		e.Description,
		e.Status.String(),
	)
	return kind
}

func mapEntities(es []*ent.Kind) []*domain.Kind {
	kinds := make([]*domain.Kind, 0, len(es))
	for _, e := range es {
		kinds = append(kinds, mapEntity(e))
	}
	return kinds
}
