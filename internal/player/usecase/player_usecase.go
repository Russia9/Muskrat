package usecase

import (
	"context"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
)

type uc struct {
	repo domain.PlayerRepository
}

func NewPlayerUsecase(repo domain.PlayerRepository) domain.PlayerUsecase {
	return &uc{repo}
}

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

func (u *uc) Get(ctx context.Context, scope permissions.Scope, id int64) (*domain.Player, error) {
	// Permission check
	if scope.PlayerRole < permissions.PlayerRoleUser {
		return nil, permissions.ErrForbidden
	}

	// Get object from repository
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "player repo")
	}

	// Permission check
	if scope.ID != obj.ID && (scope.SquadRole < permissions.SquadRoleSquire || scope.SquadID != obj.SquadID) {
		return nil, permissions.ErrForbidden
	}

	return obj, nil
}

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
	if scope.ID != obj.ID && (scope.SquadRole < permissions.SquadRoleSquire || scope.SquadID != obj.SquadID) {
		return nil, permissions.ErrForbidden
	}

	return obj, nil
}

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
		return nil, err
	}

	return obj, nil
}
