package adapters

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/kind"
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
		SetCreatedAt(k.CreatedAt()).
		SetUpdatedAt(k.UpdatedAt()).
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
	createdAt *time.Time,
	id *uuid.UUID,
	status *string,
) ([]*domain.Kind, error) {
	kindQuery := r.client.Query().
		Order(
			ent.Desc(kind.FieldCreatedAt),
			ent.Desc(kind.FieldID),
		)

	if status != nil {
		kindQuery = kindQuery.Where(
			kind.StatusEQ(kind.Status(*status)),
		)
	}
	if limit != nil {
		kindQuery = kindQuery.Limit(*limit)
	}
	switch {
	case createdAt != nil && id != nil:
		kindQuery.Where(
			func(s *sql.Selector) {
				s.Where(
					sql.CompositeLT([]string{"created_at", "id"}, *createdAt, *id),
				)
			},
		)

	case id != nil:
		kindQuery = kindQuery.Where(
			kind.IDLT(*id),
		)

	case createdAt != nil:
		kindQuery = kindQuery.Where(
			kind.CreatedAtLTE(*createdAt),
		)
	}

	es, err := kindQuery.All(ctx)
	if err != nil {
		return nil, err
	}

	return mapEntities(es), nil
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
		e.CreatedAt,
		e.UpdatedAt,
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
