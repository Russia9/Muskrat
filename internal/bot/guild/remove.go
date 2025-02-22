package guild

import (
	"context"
	"errors"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"gopkg.in/telebot.v3"
	"regexp"
	"strconv"
)

var removeGuildByTagRegex = regexp.MustCompile(`/squad_guild_remove ([a-zA-Z].+)`)
var removeGuildByLeaderIDRegex = regexp.MustCompile(`/squad_guild_remove (\d+)`)

func (m *Module) GuildRemove(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Send(m.l.Text(c, "not_in_chat"))
	}

	if removeGuildByTagRegex.MatchString(c.Text()) {
		match := removeGuildByTagRegex.FindStringSubmatch(c.Text())
		guildTag := match[1]
		err := m.guild.DeleteByTag(context.Background(), scope, guildTag, c.Chat().ID)
		if err != nil {
			if errors.Is(err, domain.ErrSquadNotFound) {
				return c.Send(m.l.Text(c, "not_in_chat"))
			} else if errors.Is(err, domain.ErrGuildNotFound) {
				return c.Send(m.l.Text(c, "guild_not_in_guild"))
			}
			return err
		}
		return c.Send(m.l.Text(c, "guild_delete_success"))
	} else if removeGuildByLeaderIDRegex.MatchString(c.Text()) {
		match := removeGuildByLeaderIDRegex.FindStringSubmatch(c.Text())
		guildLeaderID, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		}
		err = m.guild.DeleteByLeader(context.Background(), scope, int64(guildLeaderID), c.Chat().ID)
		if err != nil {
			if errors.Is(err, domain.ErrSquadNotFound) {
				return c.Send(m.l.Text(c, "not_in_chat"))
			} else if errors.Is(err, domain.ErrGuildNotFound) {
				return c.Send(m.l.Text(c, "guild_not_in_guild"))
			}
			return err
		}
		return c.Send(m.l.Text(c, "guild_delete_success"))
	}

	return c.Send(m.l.Text(c, "guild_delete_wrong_format"))
}
