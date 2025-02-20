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

func NewPlayerRepo(db *mongo.Database) domain.PlayerRepository {
	return &repo{db.Collection("players")}
}

func (r *repo) Create(ctx context.Context, obj *domain.Player) error {
	// Run InsertOne query
	_, err := r.c.InsertOne(ctx, obj)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) Get(ctx context.Context, id int64) (*domain.Player, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"_id": id})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrPlayerNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Player
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func (r *repo) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"username": username})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrPlayerNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Player
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func listAggregation(field, value string, sort domain.PlayerSort) bson.A {
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				field: value,
			},
		},
	}

	switch sort {
	case domain.PlayerSortLevel:
		pipeline = append(pipeline,
			bson.M{
				"$sort": bson.M{
					"level": -1,
				},
			})
	case domain.PlayerSortRank:
		pipeline = append(pipeline,
			bson.M{
				"$sort": bson.M{
					"rank": -1,
				},
			})
	case domain.PlayerSortBalance:
		pipeline = append(pipeline,
			bson.M{
				"$addFields": bson.M{
					"totalbalance": bson.M{
						"$add": bson.A{
							"$playerbalance",
							"$bankbalance",
						},
					},
				},
			})
		pipeline = append(pipeline,
			bson.M{
				"$sort": bson.M{
					"totalbalance": -1,
				},
			})
	}

	return pipeline
}

func (r *repo) ListBySquad(ctx context.Context, squadID string, sort domain.PlayerSort) ([]*domain.Player, error) {
	// Find objects
	cur, err := r.c.Aggregate(ctx, listAggregation("squadid", squadID, sort))
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	// Decode objects
	var objs []*domain.Player
	err = cur.All(ctx, &objs)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return objs, nil
}

func (r *repo) ListByGuild(ctx context.Context, guildID string, sort domain.PlayerSort) ([]*domain.Player, error) {
	// Find objects
	cur, err := r.c.Aggregate(ctx, listAggregation("guildid", guildID, sort))
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	// Decode objects
	var objs []*domain.Player
	err = cur.All(ctx, &objs)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return objs, nil
}

func (r *repo) CountBySquad(ctx context.Context, squadID string) (int64, error) {
	// Run CountDocuments query
	count, err := r.c.CountDocuments(ctx, bson.M{"squadid": squadID})
	if err != nil {
		return 0, errors.Wrap(err, "mongo")
	}

	return count, nil
}

func (r *repo) CountByGuild(ctx context.Context, guildID string) (int64, error) {
	// Run CountDocuments query
	count, err := r.c.CountDocuments(ctx, bson.M{"guildid": guildID})
	if err != nil {
		return 0, errors.Wrap(err, "mongo")
	}

	return count, nil
}

func (r *repo) Update(ctx context.Context, obj *domain.Player) error {
	// Run Replace query
	_, err := r.c.ReplaceOne(ctx, bson.M{"_id": obj.ID}, obj)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	// Run Delete query
	_, err := r.c.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) RemoveGuild(ctx context.Context, guildID string) error {
	//TODO implement me
	panic("implement me")
}
