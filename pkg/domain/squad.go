package domain

import (
	"context"
	"errors"
)

// Entity
type Squad struct {
	ID     string `json:"id"`
	ChatID string `json:"chat_id"`

	Name   string `json:"name"`
	Castle Castle `json:"castle"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
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
