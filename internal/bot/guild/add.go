package guild

import (
	"context"
	"errors"
	"regexp"
	"strconv"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/Russia9/Muskrat/pkg/utils"
	"gopkg.in/telebot.v3"
)

var guildRegex = regexp.MustCompile(`[🇮🇲🇻🇦🇪🇺🇲🇴]+(?:\[(.+)\] )?([\w ]*)\n.*: (.+)\n🏅Level: (\d+)`)

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

	// Check reply message format
	if !guildRegex.MatchString(c.Message().ReplyTo.Text) || c.Message().ReplyTo.Sender.ID != utils.ChatWarsBot {
		return c.Send(m.l.Text(c, "guild_add_not_reply"))
	}

	// Get Squad by chat
	sq, err := m.squad.GetByChatID(context.Background(), scope, c.Chat().ID)
	if errors.Is(err, domain.ErrSquadNotFound) {
		return c.Send(m.l.Text(c, "not_in_chat"))
	} else if err != nil {
		return err
	}
	if sq.ID != *scope.SquadID {
		return c.Send(m.l.Text(c, "not_in_chat"))
	}

	// Parse guild info
	match := guildRegex.FindStringSubmatch(c.Message().ReplyTo.Text)
	name := match[2]
	tag := match[1]
	level, _ := strconv.Atoi(match[3])

	// Create Guild
	g, err := m.guild.Create(context.Background(), scope, c.Message().ReplyTo.Sender.ID, name, tag, level)
	if errors.Is(err, domain.ErrAlreadyInGuild) {
		return c.Send(m.l.Text(c, "guild_already_in_guild"))
	} else if errors.Is(err, domain.ErrGuildAlreadyExists) {
		return c.Send(m.l.Text(c, "guild_add_already_exists", g))
	} else if err != nil {
		return err
	}

	// TODO: Success message

	return nil
}
