package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) ListBySquad(ctx context.Context, scope permissions.Scope, squadID string) ([]*domain.Guild, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadRole < permissions.SquadRoleMember || (scope.SquadID != nil && squadID != *scope.SquadID) {
		return nil, permissions.ErrForbidden
	}

	// List guilds
	gs, err := u.repo.ListBySquad(ctx, squadID)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	return gs, nil
}
