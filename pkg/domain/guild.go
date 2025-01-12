package domain

import (
	"context"
	"errors"

	"github.com/Russia9/Muskrat/pkg/permissions"
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
	Create(ctx context.Context, scope permissions.Scope, name, tag, leaderPlayerName string, level int) (*Guild, error)

	Get(ctx context.Context, scope permissions.Scope, id string) (*Guild, error)
	GetByTag(ctx context.Context, scope permissions.Scope, tag string) (*Guild, error)
	GetByLeader(ctx context.Context, scope permissions.Scope, leaderID int64) (*Guild, error)
	GetBySquadAndName(ctx context.Context, scope permissions.Scope, squadID, name string) (*Guild, error)

	ListBySquad(ctx context.Context, scope permissions.Scope, squadID string) ([]*Guild, error)

	Update(ctx context.Context, scope permissions.Scope, name, tag string, level int) (*Guild, error)
}

type GuildRepository interface {
	Create(ctx context.Context, obj *Guild) error

	Get(ctx context.Context, id string) (*Guild, error)
	GetByTag(ctx context.Context, tag string) (*Guild, error)
	GetByLeader(ctx context.Context, leaderID int64) (*Guild, error)
	GetBySquadAndName(ctx context.Context, squadID, name string) (*Guild, error)

	ListBySquad(ctx context.Context, squadID string) ([]*Guild, error)

	Update(ctx context.Context, obj *Guild) error

	Delete(ctx context.Context, id string) error
}
