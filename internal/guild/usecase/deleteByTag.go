package usecase

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u *uc) DeleteByTag(ctx context.Context, scope permissions.Scope, tag string, chatID int64) (err error) {
	if scope.SquadRole != permissions.SquadRoleLeader {
		return permissions.ErrForbidden
	}

	guild, err := u.repo.GetByTag(ctx, tag)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrGuildNotFound
		}
		return errors.Wrap(err, `guild repo`)
	}

	_, err = u.squad.GetByChatID(ctx, chatID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrSquadNotFound
		}
		return errors.Wrap(err, `squad repo`)
	}

	err = u.repo.DeleteByTag(ctx, tag)
	if err != nil {
		return errors.Wrap(err, `guild repo`)
	}

	err = u.player.RemoveGuild(ctx, guild.ID)
	if err != nil {
		return errors.Wrap(err, `player repo`)
	}

	return nil
}
