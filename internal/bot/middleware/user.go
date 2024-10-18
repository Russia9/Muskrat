package middleware

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
)

func (m *Middleware) Player(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		// Create Scope
		scope := permissions.Scope{
			ID:         c.Sender().ID,
			PlayerRole: permissions.PlayerRoleUnregistered,
		}

		// Check if player is registered
		player, err := m.player.Get(context.Background(), scope, c.Sender().ID)
		if err == nil { // Set permission level if user is found
			scope.PlayerRole = player.PlayerRole
		} else if errors.Is(err, domain.ErrPlayerNotFound) { // Register user if not found
			// Create user
			player, err = m.player.Create(context.Background(), scope, c.Sender().ID, c.Sender().Username)
			if err != nil {
				return errors.Wrap(err, "user usecase")
			}
		}

		// Check if player is banned
		if scope.PlayerRole < permissions.PlayerRoleUnregistered {
			return nil
		}

		// Run Player.Seen
		player, err = m.player.Seen(context.Background(), scope, player.Username)
		if err != nil {
			return errors.Wrap(err, "user usecase")
		}

		// Set Player and Scope in context
		c.Set("player", player)
		c.Set("scope", scope)
		m.layout.SetLocale(c, player.Language)

		// Process request
		return next(c)
	}
}
