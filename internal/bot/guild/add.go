package guild

import (
	"context"
	"errors"
	"regexp"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/Russia9/Muskrat/pkg/utils"
	"gopkg.in/telebot.v3"
)

var guildRegex = regexp.MustCompile("[🇮🇲🇻🇦🇪🇺🇲🇴]+(?:\\[(.+)\\] )?([\\w ]*)\n.*: (.+)\n🏅Level: (\\d+)")
var guildAddRegex = regexp.MustCompile("/squad_guild_add (.*)")

func (m *Module) GuildAdd(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check if chat is supergroup
	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Send(m.l.Text(c, "not_in_chat"))
	}

	// Check if message is a reply
	if c.Message().ReplyTo == nil || c.Message().ReplyTo.Sender == nil {
		return c.Send(m.l.Text(c, "squad_add_not_reply"))
	}

	// Check message format
	if !guildAddRegex.MatchString(c.Message().Text) {
		return c.Send(m.l.Text(c, "guild_add_wrong_format"))
	}

	// Check reply message format
	if !guildRegex.MatchString(c.Message().ReplyTo.Text) || c.Message().ReplyTo.Sender.ID != utils.ChatWarsBot {
		return c.Send(m.l.Text(c, "guild_add_wrong_reply_format"))
	}

	// Get squad by chat
	sq, err := m.squad.GetByChatID(context.Background(), scope, c.Chat().ID)
	if errors.Is(err, domain.ErrSquadNotFound) {
		return c.Send(m.l.Text(c, "not_in_chat"))
	} else if err != nil {
		return err
	}
	if sq.ID != *scope.SquadID {
		return c.Send(m.l.Text(c, "not_in_chat"))
	}

	// TODO: Create Guild

	return nil
}
