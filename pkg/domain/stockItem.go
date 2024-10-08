package domain

import (
	"context"
	"errors"
)

// Entity
type StockItem struct {
	PlayerID  int64  `json:"player_id"`
	Warehouse string `json:"warehouse"` // g_3_3

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
