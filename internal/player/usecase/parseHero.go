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

var heroPlayerNameRegex = regexp.MustCompile("([🇮🇲🇻🇦🇪🇺🇲🇴]+)([a-zA-Z0-9 _]{4,16})") // PlayerName & Castle
var basicStatsRegex = regexp.MustCompile("STR (\\d+) DEX (\\d+) VIT (\\d+)")

var detailedStatsRegex = map[string]*regexp.Regexp{
	"attackForce":   regexp.MustCompile("🗡️Attack Force: (\\d+)"),
	"attackSpeed":   regexp.MustCompile("🌀Attack Speed: (\\d+)"),
	"criticalRate":  regexp.MustCompile("⚡️Critical Rate: (\\d+)"),
	"criticalForce": regexp.MustCompile("💥Critical Force: (\\d+)"),
	"accuracy":      regexp.MustCompile("🦅Accuracy: (\\d+)"),
	"evasion":       regexp.MustCompile("🥋Evasion: (\\d+)"),
	"armorScore":    regexp.MustCompile("🛡️Armor Score: (\\d+)"),
	"moveSpeed":     regexp.MustCompile("🥾Move Speed: (\\d+)"),
}

// ParseHero parses the /hero message and updates the player's info
func (u *uc) ParseHero(ctx context.Context, scope permissions.Scope, hero string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Validate text
	if !heroPlayerNameRegex.MatchString(hero) {
		return nil, domain.ErrInvalidText
	}

	// Get Player from repo
	player, err := u.repo.Get(ctx, scope.ID)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Parse Castle, PlayerName
	playerName := heroPlayerNameRegex.FindStringSubmatch(hero)
	if len(playerName) != 3 {
		return nil, domain.ErrInvalidText
	}
	player.Castle = domain.FlagToCastle(playerName[1])
	player.PlayerName = playerName[2]

	// Parse Basic info
	err = parseBase(player, hero)
	if err != nil {
		return nil, err
	}

	// Parse Basic Stats
	stats := basicStatsRegex.FindStringSubmatch(hero)
	if len(stats) != 4 {
		return nil, domain.ErrInvalidText
	}
	player.Str, err = strconv.Atoi(stats[1])
	if err != nil {
		return nil, domain.ErrInvalidText
	}
	player.Dex, err = strconv.Atoi(stats[2])
	if err != nil {
		return nil, domain.ErrInvalidText
	}
	player.Vit, err = strconv.Atoi(stats[3])
	if err != nil {
		return nil, domain.ErrInvalidText
	}

	// Parse Detailed Stats
	for key, regex := range detailedStatsRegex {
		stat := regex.FindStringSubmatch(hero)
		if len(stat) != 2 {
			return nil, domain.ErrInvalidText
		}
		player.DetailedStats[key], err = strconv.Atoi(stat[1])
		if err != nil {
			return nil, domain.ErrInvalidText
		}
	}

	player.ProfileUpdatedAt = time.Now()

	// Update Player in repo
	err = u.repo.Update(ctx, player)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return player, nil
}
