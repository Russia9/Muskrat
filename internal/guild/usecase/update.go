package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Update(ctx context.Context, scope permissions.Scope, name, tag, hqLocation string, level int) (*domain.Guild, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Get Guild
	g, err := u.repo.Get(ctx, *scope.GuildID)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	// Permission check
	if scope.GuildRole < permissions.SquadRoleLeader || (scope.GuildID != nil && g.ID != *scope.GuildID) {
		return nil, permissions.ErrForbidden
	}

	// Update guild
	g.Name = name
	g.Tag = tag
	g.HQLocation = hqLocation
	g.Level = level

	err = u.repo.Update(ctx, g)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	return g, nil
}
