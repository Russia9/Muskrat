package usecase

import (
	"context"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (u *uc) Create(ctx context.Context, scope permissions.Scope, chatID int64, name string) (*domain.Squad, error) {
	// Permissions check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadID != nil || scope.SquadRole != permissions.SquadRoleNone {
		return nil, domain.ErrAlreadyInSquad
	}

	// Get player
	player, err := u.player.Get(ctx, scope.ID)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Check if player profile is updated
	if !player.Updated() {
		return nil, domain.ErrNeedProfileUpdate
	}

	// Check if chat is already attached to a squad
	_, err = u.repo.GetByChatID(ctx, chatID)
	if err == nil {
		return nil, domain.ErrChatAlreadyAttached
	} else if !errors.Is(err, domain.ErrSquadNotFound) {
		return nil, errors.Wrap(err, "squad repo")
	}

	// Create squad
	obj := &domain.Squad{
		ID:        uuid.NewString(),
		ChatID:    chatID,
		Name:      name,
		Castle:    player.Castle,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save squad
	err = u.repo.Create(ctx, obj)
	if err != nil {
		return nil, errors.Wrap(err, "squad repo")
	}

	// Update player
	player.SquadID = &obj.ID
	player.SquadRole = permissions.SquadRoleLeader

	// Save player
	err = u.player.Update(ctx, player)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return obj, nil
}
