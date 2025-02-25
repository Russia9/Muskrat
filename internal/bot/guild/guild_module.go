package guild

import (
	"github.com/Russia9/Muskrat/pkg/domain"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Module struct {
	player domain.PlayerUsecase
	squad  domain.SquadUsecase
	guild  domain.GuildUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase, squad domain.SquadUsecase, guild domain.GuildUsecase) *Module {
	m := &Module{player, squad, guild, tb, l}

	tb.Handle("/squad_guild_add", m.GuildAdd)
	tb.Handle("/squad_guild_remove", m.GuildRemove)

	tb.Handle("/guild_help", m.GuildHelp)

	return m
}
