package usecase

import (
	"context"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)


func (u *uc) ChangeName(ctx context.Context, scope permissions.Scope, name string) (*domain.Squad, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleLeader {
		return nil, permissions.ErrForbidden
	}

	// Get squad
	obj, err := u.repo.Get(ctx, *scope.SquadID)
	if err != nil {
		return nil, errors.Wrap(err, "squad repo")
	}

	// Update squad
	obj.Name = name
	obj.UpdatedAt = time.Now()

	// Save squad
	err = u.repo.Update(ctx, obj)
	if err != nil {
		return nil, errors.Wrap(err, "squad repo")
	}

	return obj, nil
}
