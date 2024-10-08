package domain

import (
	"context"
	"errors"
)

// Entity
type Guild struct {
	ID      string `json:"id"`
	SquadID string `json:"squad_id"`

	Name     string `json:"name"`
	Tag      string `json:"tag"`
	LeaderID int64  `json:"leader_id"`

	Level int `json:"level"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
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
