package usecase

import (
	"context"
	"regexp"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/Russia9/Muskrat/pkg/utils"
	"github.com/pkg/errors"
)

var schoolRegex = regexp.MustCompile(" (...) .*\n /school_(.*)")

func (u *uc) ParseSchool(ctx context.Context, scope permissions.Scope, school string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Validate text
	if !schoolRegex.MatchString(school) {
		return nil, domain.ErrInvalidText
	}

	// Get Player from repo
	player, err := u.repo.Get(ctx, scope.ID)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Parse Schools
	schools := schoolRegex.FindAllStringSubmatch(school, -1)
	player.Schools = make(map[string]int, len(schools))
	for _, school := range schools {
		if len(school) != 3 {
			return nil, domain.ErrInvalidText
		}
		player.Schools[school[2]] = utils.KeycapToDigit(school[1])
	}

	player.SchoolsUpdatedAt = time.Now()

	// Update Player in repo
	err = u.repo.Update(ctx, player)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return player, nil
}
