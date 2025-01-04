package usecase

import (
	"context"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Create(ctx context.Context, scope permissions.Scope, id int64, username string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole != permissions.PlayerRoleUnregistered {
		return nil, permissions.ErrForbidden
	}

	// Create object
	obj := &domain.Player{
		ID:         id,
		Username:   username,
		PlayerRole: permissions.PlayerRoleUser,

		SquadRole: permissions.SquadRoleNone, // Default

		Language: "ru", // Default

		FirstSeen: time.Now(),
		LastSeen:  time.Now(),
	}

	// Save to repository
	err := u.repo.Create(ctx, obj)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return obj, nil
}
