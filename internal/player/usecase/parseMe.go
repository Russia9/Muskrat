package usecase

import (
	"context"
	"regexp"
	"strconv"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

var playerNameRegex = regexp.MustCompile("([🇮🇲🇻🇦🇪🇺🇲🇴]+)([a-zA-Z0-9 _]{4,16}) explorer") // PlayerName & Castle
var levelRegex = regexp.MustCompile("🏅Level: (\\d+)")
var expRegex = regexp.MustCompile("📖Exp: (\\d+)/(\\d+)")
var rankRegex = regexp.MustCompile("⚔️Rank: (\\d+)")

var balanceRegex = regexp.MustCompile("🪙(\\d+) 💰(\\d+)")

func (u *uc) ParseMe(ctx context.Context, scope permissions.Scope, me string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Validate text
	if !playerNameRegex.MatchString(me) {
		return nil, domain.ErrInvalidText
	}

	// Get Player from repo
	player, err := u.repo.Get(ctx, scope.ID)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Parse Castle, PlayerName
	playerName := playerNameRegex.FindStringSubmatch(me)
	if len(playerName) != 3 {
		return nil, domain.ErrInvalidText
	}
	player.Castle = domain.FlagToCastle(playerName[1])
	player.PlayerName = playerName[2]

	// Parse Level
	level := levelRegex.FindStringSubmatch(me)
	if len(level) != 2 {
		return nil, domain.ErrInvalidText
	}
	player.Level, err = strconv.Atoi(level[1])
	if err != nil {
		return nil, domain.ErrInvalidText
	}

	// Parse Exp
	exp := expRegex.FindStringSubmatch(me)
	if len(exp) != 3 {
		return nil, domain.ErrInvalidText
	}
	player.CurrentExp, err = strconv.Atoi(exp[1])
	if err != nil {
		return nil, domain.ErrInvalidText
	}
	player.NextLevelExp, err = strconv.Atoi(exp[2])
	if err != nil {
		return nil, domain.ErrInvalidText
	}

	// Parse Rank
	rank := rankRegex.FindStringSubmatch(me)
	if len(exp) != 2 {
		return nil, domain.ErrInvalidText
	}
	player.Rank, err = strconv.Atoi(rank[2])

	// Parse Balance
	balance := balanceRegex.FindStringSubmatch(me)
	if len(balance) != 3 {
		return nil, domain.ErrInvalidText
	}
	player.PlayerBalance, err = strconv.Atoi(balance[1])
	if err != nil {
		return nil, domain.ErrInvalidText
	}
	player.BankBalance, err = strconv.Atoi(balance[2])
	if err != nil {
		return nil, domain.ErrInvalidText
	}

	// Update Player in repo
	err = u.repo.Update(ctx, player)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return player, nil
}
