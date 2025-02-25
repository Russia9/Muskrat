package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) DeleteByTag(ctx context.Context, scope permissions.Scope, tag string, chatID int64) (err error) {
	// Permission check
	if scope.SquadRole != permissions.SquadRoleLeader || scope.SquadID == nil {
		return permissions.ErrForbidden
	}

	// Get Guild by tag
	g, err := u.repo.GetByTag(ctx, tag)
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
