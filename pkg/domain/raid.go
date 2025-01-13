package domain

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

type Raid struct {
	Name string    `bson:"name" json:"name"`
	Cell string    `bson:"cell" json:"cell"`
	Time time.Time `bson:"time" json:"time"`
}

var ErrRaidNotFound = errors.New("raid not found")

type RaidUsecase interface {
	Create(ctx context.Context, name string, cell string, preRaidTime int32) error
	Get(ctx context.Context, raid *Raid) (*Raid, error)
}

type RaidRepository interface {
	Create(ctx context.Context, raid *Raid) error
	Get(ctx context.Context, raid *Raid) (*Raid, error)
}
