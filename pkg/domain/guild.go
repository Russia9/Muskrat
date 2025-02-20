package domain

import (
	"context"
	"errors"
	"time"

	"github.com/Russia9/Muskrat/pkg/permissions"
)

// Entity
type Guild struct {
	ID      string `bson:"_id"`
	SquadID string

	HQLocation string
	Name       string
	Tag        string
	LeaderID   int64

	Level int

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Errors
var ErrGuildNotFound = errors.New("guild not found")
var ErrGuildAlreadyExists = errors.New("guild already exists")
var ErrAlreadyInGuild = errors.New("already in guild")

// Interfaces
type GuildUsecase interface {
	Create(ctx context.Context, scope permissions.Scope, leader int64, name, tag, hqLocation string, level int) (*Guild, error)

	Get(ctx context.Context, scope permissions.Scope, id string) (*Guild, error)
	GetByTag(ctx context.Context, scope permissions.Scope, tag string) (*Guild, error)
	GetByLeader(ctx context.Context, scope permissions.Scope, leaderID int64) (*Guild, error)
	GetBySquadAndName(ctx context.Context, scope permissions.Scope, squadID, name string) (*Guild, error)

	ListBySquad(ctx context.Context, scope permissions.Scope, squadID string) ([]*Guild, error)

	Update(ctx context.Context, scope permissions.Scope, name, tag, hqLocation string, level int) (*Guild, error)

	ParseGuild(ctx context.Context, scope permissions.Scope, msg string) (*Guild, error)
	ParseList(ctx context.Context, scope permissions.Scope, msg string) (*Guild, error)
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
	DeleteByLeader(ctx context.Context, leaderID int64) error
	DeleteByTag(ctx context.Context, tag string) error
}
