package mongodb

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StockRepository interface {
	GetOne(ctx context.Context, dim, quality string) (*entities.Stock, error)
	GetAll(ctx context.Context) ([]*entities.Stock, error)

	Create(ctx context.Context, e *entities.Stock) (*entities.Stock, error)
	Update(ctx context.Context, e *entities.Stock) (*entities.Stock, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type stockRepository struct {
	Collection *mongo.Collection
}

func NewStockRepo(collection *mongo.Collection) StockRepository {
	return &stockRepository{
		Collection: collection,
	}
}

func (r *stockRepository) GetOne(ctx context.Context, dim, quality string) (*entities.Stock, error) {
	c := r.Collection.FindOne(ctx, bson.M{"dim": dim, "quality": quality})

	var stock entities.Stock
	err := c.Decode(&stock)
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *stockRepository) Create(ctx context.Context, e *entities.Stock) (*entities.Stock, error) {
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

func (r *stockRepository) GetAll(ctx context.Context) ([]*entities.Stock, error) {
	c, err := r.Collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	var stocks []*entities.Stock
	err = c.All(ctx, &stocks)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}

func (r *stockRepository) Update(ctx context.Context, e *entities.Stock) (*entities.Stock, error) {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"uuid": e.ID}, bson.M{"$set": e})
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *stockRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *stockRepository) DeleteAll(ctx context.Context) error {
	_, err := r.Collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
