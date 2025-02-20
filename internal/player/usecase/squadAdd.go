package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) SquadAdd(ctx context.Context, scope permissions.Scope, id int64) (*domain.Player, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleLeader {
		return nil, permissions.ErrForbidden
	}

	// Get pl
	pl, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "user repo")
	}

	// Check if player is already in squad
	if pl.SquadID != nil {
		return nil, domain.ErrAlreadyInSquad
	}

	// Add player to squad
	pl.SquadID = scope.SquadID
	pl.SquadRole = permissions.SquadRoleMember

	// Save player
	err = u.repo.Update(ctx, pl)
	if err != nil {
		return nil, errors.Wrap(err, "user repo")
	}

	return pl, nil
}
