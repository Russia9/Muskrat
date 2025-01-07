package usecase

import (
	"context"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

func (u *uc) ChangeChatID(ctx context.Context, scope permissions.Scope, chatID int64) (*domain.Squad, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleLeader {
		return nil, permissions.ErrForbidden
	}

	// Check if chat is already attached to a squad
	_, err := u.repo.GetByChatID(ctx, chatID)
	if err == nil {
		return nil, domain.ErrChatAlreadyAttached
	} else if !errors.Is(err, domain.ErrSquadNotFound) {
		return nil, errors.Wrap(err, "squad repo")
	}

	// Get squad
	obj, err := u.repo.Get(ctx, *scope.SquadID)
	if err != nil {
		return nil, errors.Wrap(err, "squad repo")
	}

	// Update squad
	obj.ChatID = chatID
	obj.UpdatedAt = time.Now()

	// Save squad
	err = u.repo.Update(ctx, obj)
	if err != nil {
		return nil, errors.Wrap(err, "squad repo")
	}

	return obj, nil
}
