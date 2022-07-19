package mongodb

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SpacingRepository interface {
	GetAll(ctx context.Context, machine entities.MachineType) ([]*entities.Spacing, error)
	GetOne(ctx context.Context, machine entities.MachineType, dim string, e *entities.End) (*entities.Spacing, error)
	Create(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error)
	Update(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type spacingRepository struct {
	Collection *mongo.Collection
}

func NewSpacingRepo(collection *mongo.Collection) SpacingRepository {
	return &spacingRepository{
		Collection: collection,
	}
}

func (r *spacingRepository) Create(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	e.ID = uuid

	_, err = r.Collection.InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *spacingRepository) GetAll(ctx context.Context, machine entities.MachineType) ([]*entities.Spacing, error) {
	c, err := r.Collection.Find(ctx, bson.M{"machine": machine})
	if err != nil {
		return nil, err
	}

	var spacings []*entities.Spacing
	err = c.All(ctx, &spacings)
	if err != nil {
		return nil, err
	}
	return spacings, nil
}

func (r *spacingRepository) GetOne(ctx context.Context, machine entities.MachineType, dim string, e *entities.End) (*entities.Spacing, error) {
	hasBevel := e.WebBevel != nil || e.FlangeBevel != nil
	hasScallop := e.Scallop != nil

	cursor := r.Collection.FindOne(ctx, bson.M{"machine": machine, "dim": dim, "name": e.Name, "has_bevel": hasBevel, "has_scallop": hasScallop})

	var spacing entities.Spacing
	err := cursor.Decode(&spacing)
	if err != nil {
		return nil, err
	}
	return &spacing, nil
}

func (r *spacingRepository) Update(ctx context.Context, e *entities.Spacing) (*entities.Spacing, error) {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"uuid": e.ID}, bson.M{"$set": e})
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *spacingRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *spacingRepository) DeleteAll(ctx context.Context) error {
	_, err := r.Collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
