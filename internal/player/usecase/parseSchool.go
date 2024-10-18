package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

func (u *uc) ParseSchool(ctx context.Context, scope permissions.Scope, school string) (*domain.Player, error) {
	panic("unimplemented")
}
