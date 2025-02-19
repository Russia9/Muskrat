package usecase

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) GetByTag(ctx context.Context, scope permissions.Scope, tag string) (*domain.Guild, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Get guild
	g, err := u.repo.GetByTag(ctx, tag)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	// Permission check
	if scope.SquadRole < permissions.SquadRoleMember || (scope.SquadID != nil && g.SquadID != *scope.SquadID) {
		return nil, permissions.ErrForbidden
	}

	return g, nil
}
