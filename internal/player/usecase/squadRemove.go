package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

func (u *uc) SquadRemove(ctx context.Context, scope permissions.Scope, id int64) (*domain.Player, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.ID != id {
		if scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleLeader {
			return nil, permissions.ErrForbidden
		}
	} else {
		if scope.SquadRole == permissions.SquadRoleLeader {
			return nil, domain.ErrLeaderCannotLeave
		}
	}

	// Get player
	pl, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if player is in squad
	if pl.SquadID == nil {
		return nil, domain.ErrNotInSquad
	}

	// Check if the player is in the same squad as caller
	if scope.SquadID != nil && *pl.SquadID != *scope.SquadID {
		return nil, permissions.ErrForbidden
	}

	// Remove player from squad
	pl.SquadID = nil
	pl.SquadRole = permissions.SquadRoleNone

	// Save player
	err = u.repo.Update(ctx, pl)
	if err != nil {
		return nil, err
	}

	return pl, nil
}
