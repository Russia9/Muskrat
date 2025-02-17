package usecase

import (
	"context"
	"regexp"
	"strconv"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

var mePlayerNameRegex = regexp.MustCompile(`([🇮🇲🇻🇦🇪🇺🇲🇴]+)(?:\[\w{2,3}\])?([a-zA-Z0-9 _]{4,16}) explorer`) // PlayerName & Castle

var balanceRegex = regexp.MustCompile(`🪙(\d+)(?: 💰(\d+))?`)

// ParseMe parses the Me message and updates the player's info
func (u *uc) ParseMe(ctx context.Context, scope permissions.Scope, me string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Validate text
	if !mePlayerNameRegex.MatchString(me) {
		return nil, domain.ErrInvalidText
	}

	// Get Player from repo
	player, err := u.repo.Get(ctx, scope.ID)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Parse Castle, PlayerName
	playerName := mePlayerNameRegex.FindStringSubmatch(me)
	if len(playerName) != 3 {
		return nil, domain.ErrInvalidText
	}
	player.Castle = domain.FlagToCastle(playerName[1])
	player.PlayerName = playerName[2]

	// Parse Basic info
	err = ParseBase(player, me)
	if err != nil {
		return nil, err
	}

	// Parse Balance
	balance := balanceRegex.FindStringSubmatch(me)
	if len(balance) == 3 {
		player.PlayerBalance, err = strconv.Atoi(balance[1])
		if err != nil {
			return nil, domain.ErrInvalidText
		}
		player.BankBalance, err = strconv.Atoi(balance[2])
		if err != nil {
			player.BankBalance = 0
		}
	} else {
		return nil, domain.ErrInvalidText
	}
	player.BalanceUpdatedAt = time.Now()

	// Update Player in repo
	err = u.repo.Update(ctx, player)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return player, nil
}
