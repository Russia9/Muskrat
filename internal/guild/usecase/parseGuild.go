package usecase

import (
	"context"
	"regexp"
	"strconv"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

var guildRegex = regexp.MustCompile(`[🇮🇲🇻🇦🇪🇺🇲🇴]+(?:\[(.+)\] )?([\w ]*)(?:\n📍Guild HQ: .*\[(.*)\])?[\w\W]+🏅Level: (\d+)`)

func (u *uc) ParseGuild(ctx context.Context, scope permissions.Scope, msg string) (*domain.Guild, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.GuildRole < permissions.SquadRoleLeader || scope.GuildID == nil {
		return nil, permissions.ErrForbidden
	}

	// Check input
	if !guildRegex.MatchString(msg) {
		return nil, domain.ErrInvalidText
	}

	// Get Guild
	g, err := u.repo.Get(ctx, *scope.GuildID)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	// Parse message
	match := guildRegex.FindStringSubmatch(msg)
	if len(match) != 5 {
		return nil, domain.ErrInvalidText
	}

	// Parse fields
	name := match[2]
	hq := match[3]
	level, err := strconv.Atoi(match[4])
	if err != nil {
		return nil, domain.ErrInvalidText
	}

	// Update Guild
	g.Name = name
	g.HQLocation = hq
	g.Level = level

	err = u.repo.Update(ctx, g)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	return g, nil
}
