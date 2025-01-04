package domain

import (
	"context"
	"errors"
)

type StockType string

const (
	StockTypePlayer    StockType = "player"
	StockTypeWarehouse StockType = "warehouse"
	StockTypeGuild     StockType = "guild"
)

// Entity
type StockItem struct {
	ID string `bson:"_id"`

	StockID   string // PlayerID or GuildID (Depending on StockType)
	StockType StockType
	Location  string // g_3_3 or nothing

	ItemType ItemType

	ItemID    int
	Sharpness int // If w34_1 then 1 goes here
	Quantity  int

	UpdatedAt int64
}

// Errors
var ErrStockItemNotFound = errors.New("stock item not found")

// Interfaces
type StockItemUsecase interface {
}

type StockItemRepository interface {
	Create(ctx context.Context, obj *StockItem) error
}
