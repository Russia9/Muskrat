package domain

import (
	"context"
	"errors"
	"time"

	"github.com/Russia9/Muskrat/pkg/permissions"
)

// Entity
type Squad struct {
	ID     string `bson:"_id"`
	ChatID int64

	Name   string
	Castle Castle

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Errors
var ErrSquadNotFound = errors.New("squad not found")
var ErrAlreadyInSquad = errors.New("already in squad")

// Interfaces
type SquadUsecase interface {
	Create(ctx context.Context, scope permissions.Scope, chatID int64, name string) (*Squad, error)

	Get(ctx context.Context, scope permissions.Scope, id string) (*Squad, error)
	GetByChatID(ctx context.Context, scope permissions.Scope, chatID int64) (*Squad, error)

	ChangeChatID(ctx context.Context, scope permissions.Scope, chatID int64) (*Squad, error)
	ChangeName(ctx context.Context, scope permissions.Scope, name string) (*Squad, error)

	Delete(ctx context.Context, scope permissions.Scope) error
}

type SquadRepository interface {
	Create(ctx context.Context, squad *Squad) error

	Get(ctx context.Context, id string) (*Squad, error)
	GetByChatID(ctx context.Context, chatID int64) (*Squad, error)

	Update(ctx context.Context, squad *Squad) error

	Delete(ctx context.Context, id string) error
}
