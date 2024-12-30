package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) GetByUsername(ctx context.Context, scope permissions.Scope, username string) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Username validation
	if !domain.UsernameRegex.MatchString(username) {
		return nil, domain.ErrInvalidUsername
	}

	// Get object from repository
	obj, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Permission check
	if scope.ID != obj.ID && (obj.SquadID == nil || scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleSquire || *scope.SquadID != *obj.SquadID) {
		return nil, permissions.ErrForbidden
	}

	return obj, nil
}
