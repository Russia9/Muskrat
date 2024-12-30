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
	StockID   string    `json:"stock_id"` // PlayerID or GuildID (Depending on StockType)
	StockType StockType `json:"stock_type"`
	Location  string    `json:"warehouse"` // g_3_3 or nothing

	ItemID    int `json:"item_id"`
	Sharpness int `json:"sharpness"` // If w34_1 then 1 goes here
	Quantity  int `json:"quantity"`

	UpdatedAt int64 `json:"updated_at"`
}

// Errors
var ErrStockItemNotFound = errors.New("stock item not found")

// Interfaces
type StockItemUsecase interface {
}

type StockItemRepository interface {
	Create(ctx context.Context, obj *StockItem) error
}
