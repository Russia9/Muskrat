package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

func (u *uc) SquadAdd(ctx context.Context, scope permissions.Scope, id int64) (*domain.Player, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleLeader {
		return nil, permissions.ErrForbidden
	}

	// Get player
	player, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if player is already in squad
	if player.SquadID != nil {
		return nil, domain.ErrAlreadyInSquad
	}

	// Add player to squad
	player.SquadID = scope.SquadID

	// Save player
	err = u.repo.Update(ctx, player)
	if err != nil {
		return nil, err
	}

	return player, nil
}
