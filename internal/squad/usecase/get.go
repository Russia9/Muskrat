package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Get(ctx context.Context, scope permissions.Scope, id string) (*domain.Squad, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.PlayerRole < permissions.PlayerRoleRoot && (scope.SquadID == nil || *scope.SquadID != id) {
		return nil, permissions.ErrForbidden
	}

	// Get squad
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "squad repo")
	}

	return obj, nil
}
