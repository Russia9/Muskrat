package usecase

import (
	"context"
	"regexp"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

var playerNameRegex = regexp.MustCompile("([🇮🇲🇻🇦🇪🇺🇲🇴]+)([a-zA-Z0-9 _]{4,16}) explorer") // PlayerName & Castle
var levelRegex = regexp.MustCompile("🏅Level: (\\d+)")
var expRegex = regexp.MustCompile("📖Exp: (\\d+)/(\\d+)")
var rankRegex = regexp.MustCompile("⚔️Rank: (\\d+)")

var balanceRegex = regexp.MustCompile("🪙(\\d+) 💰(\\d+)")

// ParseMe implements domain.PlayerUsecase.
func (u *uc) ParseMe(ctx context.Context, scope permissions.Scope, me string) (*domain.Player, error) {
	panic("unimplemented")
}
