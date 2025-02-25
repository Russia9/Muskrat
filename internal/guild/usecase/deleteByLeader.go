package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) DeleteByLeader(ctx context.Context, scope permissions.Scope, leaderID int64, chatID int64) error {
	// Permission check
	if scope.SquadRole != permissions.SquadRoleLeader || scope.SquadID == nil {
		return permissions.ErrForbidden
	}

	// Get Guild by Leader
	g, err := u.repo.GetByLeader(ctx, leaderID)
	if err != nil {
		return errors.Wrap(err, `guild repo`)
	}

	// Check if guild is in the caller's squad
	if *scope.SquadID != g.SquadID {
		return permissions.ErrForbidden
	}

	// Get Squad by ChatID and check if guild is in the squad
	sq, err := u.squad.GetByChatID(ctx, chatID)
	if err != nil {
		return errors.Wrap(err, `squad repo`)
	}
	if sq.ID != g.SquadID {
		return domain.ErrNotInSquadChat
	}

	// Remove players from the guild
	err = u.player.RemoveGuild(ctx, g.ID)
	if err != nil {
		return errors.Wrap(err, `player repo`)
	}

	// Delete guild
	err = u.repo.Delete(ctx, g.ID)
	if err != nil {
		return errors.Wrap(err, `guild repo`)
	}

	return nil
}
