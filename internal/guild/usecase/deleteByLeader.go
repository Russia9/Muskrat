package usecase

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

func (u *uc) DeleteByLeader(ctx context.Context, scope permissions.Scope, leaderID int64) error {
	return nil
}
