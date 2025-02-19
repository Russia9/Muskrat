package usecase

import (
	"context"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (u *uc) Create(ctx context.Context, scope permissions.Scope, leader int64, name, tag, hqLocation string, level int) (*domain.Guild, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}
	if scope.SquadID == nil || scope.SquadRole < permissions.SquadRoleLeader {
		return nil, permissions.ErrForbidden
	}

	// Get leader
	pl, err := u.player.Get(ctx, leader)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Check if player is in squad
	if pl.SquadID == nil || *pl.SquadID != *scope.SquadID {
		return nil, permissions.ErrForbidden
	}

	// Check if player is already in a guild
	if pl.GuildID != nil {
		return nil, domain.ErrAlreadyInGuild
	}

	// Check if guild with the same name already exists
	_, err = u.repo.GetBySquadAndName(ctx, *scope.SquadID, name)
	if err == nil {
		return nil, domain.ErrGuildAlreadyExists
	}

	// Create guild
	g := &domain.Guild{
		ID:         uuid.NewString(),
		SquadID:    *scope.SquadID,
		Name:       name,
		Tag:        tag,
		HQLocation: hqLocation,
		LeaderID:   leader,
		Level:      level,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save guild
	err = u.repo.Create(ctx, g)
	if err != nil {
		return nil, errors.Wrap(err, "guild repo")
	}

	// Update player
	pl.GuildID = &g.ID

	// Save player
	err = u.player.Update(ctx, pl)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	return g, nil
}
