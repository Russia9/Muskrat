package domain

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

// Entity
type Raid struct {
	Name string    `bson:"name" json:"name"`
	Cell string    `bson:"cell" json:"cell"`
	Time time.Time `bson:"time" json:"time"`
}

// Errors
var ErrRaidNotFound = errors.New("raid not found")

// Interfaces
type RaidUsecase interface {
	UpdateOrCreate(ctx context.Context, name string, cell string, preRaidTime int32) error
	List(ctx context.Context) (*Raid, error)
}

type RaidRepository interface {
	Create(ctx context.Context, raid *Raid) error
	List(ctx context.Context) (*Raid, error)
}
