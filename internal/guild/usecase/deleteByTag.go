package usecase

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) DeleteByTag(ctx context.Context, scope permissions.Scope, tag string) (err error) {
	if scope.SquadRole != permissions.SquadRoleLeader {
		return permissions.ErrForbidden
	}

	err = u.repo.DeleteByTag(ctx, tag)
	if err != nil {
		return errors.Wrap(err, `guild repo`)
	}
	return nil
}
