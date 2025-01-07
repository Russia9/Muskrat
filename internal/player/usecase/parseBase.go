package usecase

import (
	"regexp"
	"strconv"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
)

var levelRegex = regexp.MustCompile("🏅Level: (\\d+)")
var expRegex = regexp.MustCompile("📖Exp: (\\d+)/(\\d+)")
var rankRegex = regexp.MustCompile("⚔️Rank: (\\d+)")

// ParseBase parses the basic info of a player:
// - Level
// - Exp
// - Rank
// Format of these fields is consistent between /hero and me
func ParseBase(player *domain.Player, msg string) error {
	var err error

	// Parse Level
	level := levelRegex.FindStringSubmatch(msg)
	if len(level) != 2 {
		return domain.ErrInvalidText
	}
	player.Level, err = strconv.Atoi(level[1])
	if err != nil {
		return domain.ErrInvalidText
	}

	// Parse Exp
	exp := expRegex.FindStringSubmatch(msg)
	if len(exp) != 3 {
		return domain.ErrInvalidText
	}
	player.CurrentExp, err = strconv.Atoi(exp[1])
	if err != nil {
		return domain.ErrInvalidText
	}
	player.NextLevelExp, err = strconv.Atoi(exp[2])
	if err != nil {
		return domain.ErrInvalidText
	}

	// Parse Rank
	rank := rankRegex.FindStringSubmatch(msg)
	if len(rank) != 2 {
		return domain.ErrInvalidText
	}
	player.Rank, err = strconv.Atoi(rank[1])

	// Update time
	player.BasicsUpdatedAt = time.Now()

	return nil
}
