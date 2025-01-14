package repository

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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

func (r repo) List(ctx context.Context) ([]*domain.Raid, error) {
	var res []*domain.Raid
	cur, err := r.c.Find(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var obj domain.Raid
		err := cur.Decode(&obj)
		if err != nil {
			log.Error().Msg("Failed to decode Raid object")
		}
	}
	return res, nil
}
