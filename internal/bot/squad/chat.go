package squad

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
)

func (m *Module) SquadChat(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check if chat is supergroup
	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Send(m.l.Text(c, "not_in_chat"))
	}

	// Check if caller is a chat admin
	cm, err := m.tb.ChatMemberOf(c.Chat(), c.Sender())
	if err != nil {
		return errors.Wrap(err, "telebot")
	}
	if cm.Role != telebot.Administrator && cm.Role != telebot.Creator {
		return c.Send(m.l.Text(c, "not_chat_admin"))
	}

	sq, err := m.squad.ChangeChatID(context.Background(), scope, c.Chat().ID)
	if errors.Is(err, domain.ErrNotInSquad) {
		return c.Send(m.l.Text(c, "squad_not_in_squad"))
	} else if errors.Is(err, domain.ErrChatAlreadyAttached) {
		return c.Send(m.l.Text(c, "squad_chat_already_attached"))
	} else if err != nil {
		return errors.Wrap(err, "squad uc")
	}

	return c.Send(m.l.Text(c, "squad_chat_success", sq))
}
