package mongo

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	c *mongo.Collection
}

func NewSquadRepo(db *mongo.Database) domain.SquadRepository {
	return &repo{db.Collection("squads")}
}

func (r *repo) Create(ctx context.Context, obj *domain.Squad) error {
	// Run InsertOne query
	_, err := r.c.InsertOne(ctx, obj)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	// Run Delete query
	_, err := r.c.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) Get(ctx context.Context, id string) (*domain.Squad, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"_id": id})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrPlayerNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Squad
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func (r *repo) GetByChatID(ctx context.Context, chatID int64) (*domain.Squad, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"chatid": chatID})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrPlayerNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Squad
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func (r *repo) Update(ctx context.Context, squad *domain.Squad) error {
	panic("unimplemented")
}
