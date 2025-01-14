package repository

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

func NewRaidRepo(db *mongo.Database) domain.RaidRepository {
	return &repo{c: db.Collection("raids")}
}

func (r repo) UpdateOrCreate(ctx context.Context, obj *domain.Raid) error {
	filter := bson.D{{
		"$and",
		bson.A{
			bson.D{
				{"name", obj.Name},
			},
			bson.D{
				{"cell", obj.Cell},
			},
		},
	}}
	res, err := r.c.ReplaceOne(ctx, filter, obj)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	if res.MatchedCount == 0 {
		_, err := r.c.InsertOne(ctx, obj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r repo) List(ctx context.Context) (*domain.Raid, error) {
	//TODO implement me
	panic("implement me")
}
