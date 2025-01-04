package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) GetByChatID(ctx context.Context, scope permissions.Scope, chatID int64) (*domain.Squad, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.PlayerRole < permissions.PlayerRoleRoot && scope.SquadID == nil {
		return nil, permissions.ErrForbidden
	}

	// Get squad
	obj, err := u.repo.GetByChatID(ctx, chatID)
	if err != nil {
		return nil, errors.Wrap(err, "squad repo")
	}

	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleRoot && *scope.SquadID != obj.ID {
		return nil, permissions.ErrForbidden
	}

	return obj, nil
}
