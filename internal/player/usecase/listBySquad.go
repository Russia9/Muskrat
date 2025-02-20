package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) ListBySquad(ctx context.Context, scope permissions.Scope, squadID string, sort domain.PlayerSort) ([]*domain.Player, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadRole < permissions.SquadRoleMember || scope.SquadID == nil {
		return nil, permissions.ErrForbidden
	}
	if scope.PlayerRole < permissions.PlayerRoleInternal && *scope.SquadID != squadID {
		return nil, permissions.ErrForbidden
	}

	// Fetch players
	players, err := u.repo.ListBySquad(ctx, squadID, sort)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return players, nil
}
