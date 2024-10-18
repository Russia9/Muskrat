package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

func (u *uc) ParseHero(ctx context.Context, scope permissions.Scope, hero string) (*domain.Player, error) {
	panic("unimplemented")
}
