package mongodb

import (
	"context"

	"github.com/edanko/moses/internal/entities"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProfileRepository interface {
	GetOne(ctx context.Context, id uuid.UUID) (*entities.Profile, error)
	GetAll(ctx context.Context) ([]*entities.Profile, error)
	GetAllLaunch(ctx context.Context, launch string) ([]*entities.Profile, error)
	Get(ctx context.Context, project, dimension, quality string) ([]*entities.Profile, error)

	Create(ctx context.Context, e *entities.Profile) (*entities.Profile, error)
	Update(ctx context.Context, e *entities.Profile) (*entities.Profile, error)
	Delete(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type profileRepository struct {
	Collection *mongo.Collection
}

func NewProfileRepo(collection *mongo.Collection) ProfileRepository {
	return &profileRepository{
		Collection: collection,
	}
}

func (p *profileRepository) GetOne(ctx context.Context, id uuid.UUID) (*entities.Profile, error) {
	c := p.Collection.FindOne(ctx, bson.M{"uuid": id})

	var profile entities.Profile
	err := c.Decode(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (p *profileRepository) Create(ctx context.Context, e *entities.Profile) (*entities.Profile, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	e.ID = uuid

	_, err = p.Collection.InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (p *profileRepository) GetAll(ctx context.Context) ([]*entities.Profile, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"name": 1})

	c, err := p.Collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	var profiles []*entities.Profile
	err = c.All(ctx, &profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p *profileRepository) GetAllLaunch(ctx context.Context, launch string) ([]*entities.Profile, error) {
	c, err := p.Collection.Find(ctx, bson.M{"launch": launch})
	if err != nil {
		return nil, err
	}

	var profiles []*entities.Profile
	err = c.All(ctx, &profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p *profileRepository) Get(ctx context.Context, project, dim, quality string) ([]*entities.Profile, error) {
	c, err := p.Collection.Find(ctx, bson.M{"project": project, "dim": dim, "quality": quality})
	if err != nil {
		return nil, err
	}

	var profiles []*entities.Profile
	err = c.All(ctx, &profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p *profileRepository) Update(ctx context.Context, e *entities.Profile) (*entities.Profile, error) {
	_, err := p.Collection.UpdateOne(ctx, bson.M{"uuid": e.ID}, bson.M{"$set": e})
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (p *profileRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := p.Collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return err
	}
	return nil
}

func (p *profileRepository) DeleteAll(ctx context.Context) error {
	_, err := p.Collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}
