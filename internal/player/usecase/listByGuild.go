package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) ListByGuild(ctx context.Context, scope permissions.Scope, guildID string, sort domain.PlayerSort) ([]*domain.Player, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.GuildRole < permissions.SquadRoleMember || scope.GuildID == nil {
		return nil, permissions.ErrForbidden
	}
	if scope.PlayerRole < permissions.PlayerRoleInternal && *scope.GuildID != guildID {
		return nil, permissions.ErrForbidden
	}

	// Fetch players
	players, err := u.repo.ListByGuild(ctx, guildID, sort)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return players, nil
}
