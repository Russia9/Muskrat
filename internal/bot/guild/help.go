package guild

import "gopkg.in/telebot.v3"

func (m Module) GuildHelp(c telebot.Context) error {
	return c.Send(m.l.Text(c, "guild_help"))
}
