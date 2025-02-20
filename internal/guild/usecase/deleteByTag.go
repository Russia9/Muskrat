package usecase

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

func (u *uc) DeleteByTag(ctx context.Context, scope permissions.Scope, tag string) (err error) {
	return nil
}
