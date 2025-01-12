package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Locale(ctx context.Context, scope permissions.Scope, lang string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Check if language is valid
	valid := false
	for _, l := range domain.Languages {
		if l == lang {
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
	pl.Locale = lang

	// Save player
	err = u.repo.Update(ctx, pl)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return pl, nil
}
