package domain

import (
	"context"
	"errors"
)

// Entity
type Guild struct {
	ID      string `bson:"_id"`
	SquadID string

	Name     string
	Tag      string
	LeaderID int64

	Level int

	CreatedAt int64
	UpdatedAt int64
}

// Errors
var ErrGuildNotFound = errors.New("guild not found")

// Interfaces
type GuildUsecase interface {
}

type GuildRepository interface {
	Create(ctx context.Context, obj *Guild) error

	Get(ctx context.Context, id string) (*Guild, error)
	GetByTag(ctx context.Context, tag string) (*Guild, error)

	ListBySquad(ctx context.Context, squadID string) ([]*Guild, error)

	Update(ctx context.Context, obj *Guild) error

	Delete(ctx context.Context, id string) error
}
