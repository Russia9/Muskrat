package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

func (u *uc) ParseList(ctx context.Context, scope permissions.Scope, idlist string) (*domain.Guild, error) {
	panic("unimplemented")
}
