package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Locale(ctx context.Context, scope permissions.Scope, locale string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Check if locale is valid
	valid := false
	for _, l := range domain.Locales {
		if l == locale {
			valid = true
			break
		}
	}
	if !valid {
		return nil, domain.ErrUnsupportedLanguage
	}

	// Get pl
	pl, err := u.repo.Get(ctx, scope.ID)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Update language
	pl.Locale = locale

	// Save player
	err = u.repo.Update(ctx, pl)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return pl, nil
}
