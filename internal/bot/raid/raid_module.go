package raid

import (
	"github.com/Russia9/Muskrat/pkg/domain"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Module struct {
	raid domain.RaidUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, raid domain.RaidUsecase) *Module {
	m := &Module{raid: raid, tb: tb, l: l}

	tb.Handle("/raids", m.List)

	return m
}
