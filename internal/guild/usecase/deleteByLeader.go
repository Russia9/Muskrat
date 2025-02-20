package usecase

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) DeleteByLeader(ctx context.Context, scope permissions.Scope, leaderID int64) error {
	if scope.SquadRole != permissions.SquadRoleLeader {
		return permissions.ErrForbidden
	}

	guild, err := u.repo.GetByLeader(ctx, leaderID)
	if err != nil {
		return errors.Wrap(err, `guild repo`)
	}

	err = u.repo.DeleteByLeader(ctx, leaderID)
	if err != nil {
		return errors.Wrap(err, `guild repo`)
	}

	err = u.player.RemoveGuild(ctx, guild.ID)
	if err != nil {
		return errors.Wrap(err, `player repo`)
	}
	return nil
}
