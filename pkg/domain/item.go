package domain

import (
	"context"
	"errors"
)

// Entity
type Item struct {
	ID   string
	Name string
	Type ItemType
}

// Constants
type ItemType string

const (
	ItemTypeResources   ItemType = "resources"
	ItemTypeEquipment   ItemType = "equipment"
	ItemTypeCraft       ItemType = "craft"
	ItemTypeFood        ItemType = "food"
	ItemTypeConsumables ItemType = "consumables"
	ItemTypeMisc        ItemType = "misc"
)

// Errors
var ErrItemNotFound = errors.New("item not found")

// Interfaces
type ItemUsecase interface {
}

type ItemTypeRepository interface {
	// Items are created by migrations

	Get(ctx context.Context, id string) (*Item, error)
	GetByName(ctx context.Context, name string) (*Item, error)

	List(ctx context.Context) ([]*Item, error)
	ListByType(ctx context.Context, itemType ItemType) ([]*Item, error)
}
