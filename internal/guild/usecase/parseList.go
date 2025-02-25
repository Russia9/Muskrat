package usecase

import (
	"context"
	"regexp"
	"strconv"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

var memberRegex = regexp.MustCompile(`👣\d+ (\d+)`)

func (u *uc) ParseList(ctx context.Context, scope permissions.Scope, idlist string) (*domain.Guild, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.GuildRole < permissions.SquadRoleLeader || scope.GuildID == nil {
		return nil, permissions.ErrForbidden
	}

	// Check input
	if !memberRegex.MatchString(idlist) {
		return nil, domain.ErrInvalidText
	}

	// Get Guild
	g, err := u.repo.Get(ctx, *scope.GuildID)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	// Parse list
	ids := make([]int64, 0)
	for _, match := range memberRegex.FindAllStringSubmatch(idlist, -1) {
		id, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			return nil, domain.ErrInvalidText
		}
		ids = append(ids, id)
	}

	// Get current members
	members, err := u.player.ListByGuild(ctx, g.ID, domain.PlayerSortRank)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Remove players that are not in the list
	for _, member := range members {
		remove := true
		for i, id := range ids {
			if member.ID == id {
				// Remove id
				ids = append(ids[:i], ids[i+1:]...)
				remove = false
				break
			}
		}

		// Remove member from guild
		if remove {
			member.GuildRole = permissions.SquadRoleNone
			member.GuildID = nil

			err := u.player.Update(ctx, member)
			if err != nil {
				return nil, errors.Wrap(err, "player repo")
			}
		}
	}

	// Add players that are in the list but not in the guild
	for _, id := range ids {
		player, err := u.player.Get(ctx, id)
		if err != nil {
			continue // Skip if player not found
		}

		player.GuildID = &g.ID
		player.GuildRole = permissions.SquadRoleMember

		err = u.player.Update(ctx, player)
		if err != nil {
			return nil, errors.Wrap(err, "player repo")
		}
	}

	return g, nil
}
