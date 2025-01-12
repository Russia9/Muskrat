package squad

import (
	"github.com/Russia9/Muskrat/pkg/domain"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Module struct {
	player domain.PlayerUsecase
	squad  domain.SquadUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase, squad domain.SquadUsecase) *Module {
	m := &Module{player, squad, tb, l}

	tb.Handle("/squad_create", m.SquadCreate)
	tb.Handle("/squad_rename", m.SquadRename)
	tb.Handle("/squad_chat", m.SquadChat)

	tb.Handle("/squad_help", m.SquadHelp)

	tb.Handle("/squad_add", m.SquadAdd)
	tb.Handle("/squad_kick", m.SquadKick)

	tb.Handle("/squad_delete", m.SquadDelete)
	tb.Handle("/squad_delete_confirm", m.SquadDeleteConfirm)

	return m
}
