package domain

import (
	"context"
	"errors"
)

// Entity
type Squad struct {
	ID     string `bson:"_id"`
	ChatID string

	Name   string
	Castle Castle

	CreatedAt int64
	UpdatedAt int64
}

// Errors
var ErrSquadNotFound = errors.New("squad not found")

// Interfaces
type SquadUsecase interface {
}

type SquadRepository interface {
	Create(ctx context.Context, squad *Squad) error

	Get(ctx context.Context, id string) (*Squad, error)
	GetByChatID(ctx context.Context, chatID string) (*Squad, error)

	Update(ctx context.Context, squad *Squad) error

	Delete(ctx context.Context, id string) error
}
