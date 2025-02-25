package usecase

import (
	"context"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Seen(ctx context.Context, scope permissions.Scope, username string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Get object from repository
	obj, err := u.repo.Get(ctx, scope.ID)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Update object
	obj.Username = username
	obj.LastSeen = time.Now()

	// Save object to repository
	err = u.repo.Update(ctx, obj)
	if err != nil {
		return nil, errors.Wrap(err, "user repo")
	}

	return obj, nil
}
