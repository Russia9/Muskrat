package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) Get(ctx context.Context, scope permissions.Scope, id int64) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUnregistered {
		return nil, permissions.ErrForbidden
	}
	if scope.ID != id { // TODO: Add permission for viewing squad members for squad leaders
		return nil, permissions.ErrForbidden
	}

	// Get object from repository
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Permission check
	if scope.ID != obj.ID && (obj.SquadID == nil || scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleSquire || *scope.SquadID != *obj.SquadID) {
		return nil, permissions.ErrForbidden
	}

	return obj, nil
}
