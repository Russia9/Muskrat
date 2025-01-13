package repository

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	c *mongo.Collection
}

func NewRaidRepo(db *mongo.Database) domain.RaidRepository {
	return &repo{c: db.Collection("raids")}
}

func (r repo) Create(ctx context.Context, obj *domain.Raid) error {
	_, err := r.c.InsertOne(ctx, obj)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r repo) Get(ctx context.Context, raid *domain.Raid) (*domain.Raid, error) {
	//TODO implement me
	panic("implement me")
}
