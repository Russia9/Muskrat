package domain

import (
	"context"
	"errors"
)

// Entity
type Recipe struct {
	ItemID string

	RequiredSchool      string
	RequiredSchoolLevel int

	Mana       int
	Components map[string]int // ItemID -> Quantity
}

// Errors
var ErrRecipeNotFound = errors.New("recipe not found")

// Interfaces
type RecipeUsecase interface {
}

type RecipeRepository interface {
	// Recipes are created by migrations

	Get(ctx context.Context, itemID string) (*Recipe, error)
	GetByItemID(ctx context.Context, itemID string) (*Recipe, error)

	List(ctx context.Context) ([]*Recipe, error)
}
