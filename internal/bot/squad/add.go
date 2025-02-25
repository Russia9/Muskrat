package squad

import (
	"context"
	"errors"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"gopkg.in/telebot.v3"
)

func (m *Module) SquadAdd(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check if chat is supergroup
	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Send(m.l.Text(c, "not_in_chat"))
	}

	// Check if message is a reply
	if c.Message().ReplyTo == nil || c.Message().ReplyTo.Sender == nil {
		return c.Send(m.l.Text(c, "squad_add_not_reply"))
	}

	// Get squad by chat
	sq, err := m.squad.GetByChatID(context.Background(), scope, c.Chat().ID)
	if errors.Is(err, domain.ErrSquadNotFound) {
		return c.Send(m.l.Text(c, "not_in_chat"))
	} else if err != nil {
		return err
	}

	// Check if caller is member of the same squad
	if scope.SquadID == nil || sq.ID != *scope.SquadID {
		return c.Send(m.l.Text(c, "squad_wrong_squad", sq))
	}

	// Add player to squad
	pl, err := m.player.SquadAdd(context.Background(), scope, c.Message().ReplyTo.Sender.ID)
	if errors.Is(err, domain.ErrAlreadyInSquad) {
		return c.Send(m.l.Text(c, "squad_add_already_in_squad", pl))
	} else if err != nil {
		return err
	}

	return c.Send(m.l.Text(c, "squad_add_success", pl))
}
