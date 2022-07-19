package mongodb

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RemnantRepository interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Remnant, error)
	GetAll(ctx context.Context) ([]*entities.Remnant, error)
	GetNotUsed(ctx context.Context, project, dim, quality string) ([]*entities.Remnant, error)
	GetAllNotUsed(ctx context.Context) ([]*entities.Remnant, error)

	Create(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error)
	Update(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type remnantRepository struct {
	Collection *mongo.Collection
}

func NewRemnantRepo(collection *mongo.Collection) RemnantRepository {
	return &remnantRepository{
		Collection: collection,
	}
}

func (r *remnantRepository) GetOne(ctx context.Context, id uuid.UUID) (*entities.Remnant, error) {
	c := r.Collection.FindOne(ctx, bson.M{"uuid": id})

	var remnant entities.Remnant
	err := c.Decode(&remnant)
	if err != nil {
		return nil, err
	}
	return &remnant, nil
}

func (r *remnantRepository) Create(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error) {
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

func (r *remnantRepository) GetAll(ctx context.Context) ([]*entities.Remnant, error) {
	c, err := r.Collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	var remnants []*entities.Remnant
	err = c.All(ctx, &remnants)
	if err != nil {
		return nil, err
	}
	return remnants, nil
}

func (r *remnantRepository) GetNotUsed(ctx context.Context, project, dim, quality string) ([]*entities.Remnant, error) {
	c, err := r.Collection.Find(ctx, bson.M{"project": project, "dim": dim, "quality": quality, "used": false})
	if err != nil {
		return nil, err
	}

	var remnants []*entities.Remnant
	err = c.All(ctx, &remnants)
	if err != nil {
		return nil, err
	}
	return remnants, nil
}

func (r *remnantRepository) GetAllNotUsed(ctx context.Context) ([]*entities.Remnant, error) {
	c, err := r.Collection.Find(ctx, bson.M{"used": false})
	if err != nil {
		return nil, err
	}

	var remnants []*entities.Remnant
	err = c.All(ctx, &remnants)
	if err != nil {
		return nil, err
	}
	return remnants, nil
}

func (r *remnantRepository) Update(ctx context.Context, e *entities.Remnant) (*entities.Remnant, error) {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"uuid": e.ID}, bson.M{"$set": e})
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *remnantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *remnantRepository) DeleteAll(ctx context.Context) error {
	_, err := r.Collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
