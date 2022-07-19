package mongodb

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NestRepository interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Nest, error)
	GetAll(ctx context.Context) ([]*entities.Nest, error)
	Get(ctx context.Context, project, dimension, quality string) ([]*entities.Nest, error)

	Create(ctx context.Context, e *entities.Nest) (*entities.Nest, error)
	Update(ctx context.Context, e *entities.Nest) (*entities.Nest, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type nestRepository struct {
	Collection *mongo.Collection
}

func NewNestRepo(collection *mongo.Collection) NestRepository {
	return &nestRepository{
		Collection: collection,
	}
}

func (n *nestRepository) GetOne(ctx context.Context, id uuid.UUID) (*entities.Nest, error) {
	c := n.Collection.FindOne(ctx, bson.M{"uuid": id})

	var nest entities.Nest
	err := c.Decode(&nest)
	if err != nil {
		return nil, err
	}
	return &nest, nil
}

func (n *nestRepository) Create(ctx context.Context, e *entities.Nest) (*entities.Nest, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	e.ID = uuid

	_, err = n.Collection.InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (n *nestRepository) GetAll(ctx context.Context) ([]*entities.Nest, error) {
	c, err := n.Collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	var nests []*entities.Nest
	err = c.All(ctx, &nests)
	if err != nil {
		return nil, err
	}

	return nests, nil
}

func (n *nestRepository) Get(ctx context.Context, project, dimension, quality string) ([]*entities.Nest, error) {
	c, err := n.Collection.Find(ctx, bson.M{"project": project, "dimension": dimension, "quality": quality})
	if err != nil {
		return nil, err
	}

	var nests []*entities.Nest
	err = c.All(ctx, &nests)
	if err != nil {
		return nil, err
	}
	return nests, nil
}

func (n *nestRepository) Update(ctx context.Context, e *entities.Nest) (*entities.Nest, error) {
	_, err := n.Collection.UpdateOne(ctx, bson.M{"uuid": e.ID}, bson.M{"$set": e})
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (n *nestRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := n.Collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return err
	}
	return nil
}

func (n *nestRepository) DeleteAll(ctx context.Context) error {
	_, err := n.Collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
