package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Delete(ctx context.Context, scope permissions.Scope) error {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return permissions.ErrForbidden
	}
	if scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleLeader {
		return permissions.ErrForbidden
	}

	// Get all squad members
	members, err := u.player.ListBySquad(ctx, *scope.SquadID)
	if err != nil {
		return errors.Wrap(err, "player repo")
	}

	// Remove all members from squad
	for _, member := range members {
		member.SquadID = nil
		member.GuildID = nil
		member.SquadRole = permissions.SquadRoleNone

		err = u.player.Update(ctx, member)
		if err != nil {
			return errors.Wrap(err, "player repo")
		}
	}

	// Remove squad
	err = u.repo.Delete(ctx, *scope.SquadID)
	if err != nil {
		return errors.Wrap(err, "squad repo")
	}

	return nil
}
