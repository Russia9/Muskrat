package mongo

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/permissions"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	c *mongo.Collection
}

func NewGuildRepo(db *mongo.Database) domain.GuildRepository {
	return &repo{db.Collection("guilds")}
}

func (r *repo) Create(ctx context.Context, obj *domain.Guild) error {
	// Insert object
	_, err := r.c.InsertOne(ctx, obj)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	// Delete object
	_, err := r.c.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) DeleteByLeader(ctx context.Context, leaderID int64) error {
	_, err := r.c.DeleteOne(ctx, bson.M{"leaderid": leaderID})
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) DeleteByTag(ctx context.Context, scope permissions.Scope, tag string) error {
	_, err := r.c.DeleteOne(ctx, bson.M{"tag": tag})
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}

func (r *repo) Get(ctx context.Context, id string) (*domain.Guild, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"_id": id})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrSquadNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Guild
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func (r *repo) GetByLeader(ctx context.Context, leaderID int64) (*domain.Guild, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"leaderid": leaderID})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrSquadNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Guild
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func (r *repo) GetBySquadAndName(ctx context.Context, squadID string, name string) (*domain.Guild, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"squadid": squadID, "name": name})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrSquadNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Guild
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func (r *repo) GetByTag(ctx context.Context, tag string) (*domain.Guild, error) {
	// Find object
	res := r.c.FindOne(ctx, bson.M{"tag": tag})
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, domain.ErrSquadNotFound
	} else if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "mongo")
	}

	// Decode object
	var obj domain.Guild
	err := res.Decode(&obj)
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}

	return &obj, nil
}

func (r *repo) ListBySquad(ctx context.Context, squadID string) ([]*domain.Guild, error) {
	// Find objects
	cur, err := r.c.Find(ctx, bson.M{"squadid": squadID})
	if err != nil {
		return nil, errors.Wrap(err, "mongo")
	}
	defer cur.Close(ctx)

	// Decode objects
	var objs []*domain.Guild
	for cur.Next(ctx) {
		var obj domain.Guild
		err := cur.Decode(&obj)
		if err != nil {
			return nil, errors.Wrap(err, "mongo")
		}

		objs = append(objs, &obj)
	}

	return objs, nil
}

func (r *repo) Update(ctx context.Context, squad *domain.Guild) error {
	// Update object
	_, err := r.c.ReplaceOne(ctx, bson.M{"_id": squad.ID}, squad)
	if err != nil {
		return errors.Wrap(err, "mongo")
	}

	return nil
}
