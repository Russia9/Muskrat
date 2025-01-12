package bot

import (
	"github.com/Russia9/Muskrat/internal/bot/finance"
	"github.com/Russia9/Muskrat/internal/bot/middleware"
	"github.com/Russia9/Muskrat/internal/bot/parse"
	"github.com/Russia9/Muskrat/internal/bot/settings"
	"github.com/Russia9/Muskrat/internal/bot/squad"
	"github.com/Russia9/Muskrat/pkg/domain"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Bot struct {
	tb *telebot.Bot
	l  *layout.Layout

	settings *settings.Module
	parse    *parse.Module
	squad    *squad.Module

	finance *finance.Module
}

func NewBot(tb *telebot.Bot, l *layout.Layout, pl domain.PlayerUsecase, sq domain.SquadUsecase) *Bot {
	b := &Bot{
		tb: tb,
		l:  l,
	}

	// Register middleware
	m := middleware.NewMiddleware(pl, l)
	b.tb.Use(m.Logger)
	b.tb.Use(m.Player)

	// Create Modules
	b.settings = settings.NewModule(tb, l, pl)
	b.parse = parse.NewModule(tb, l, pl)
	b.squad = squad.NewModule(tb, l, pl, sq)

	b.finance = finance.NewModule(tb, l, pl, sq)

	// Register handlers
	b.tb.Handle(telebot.OnText, b.Router)

	return b
}

func (b *Bot) StartAsync() {
	go b.StartBlocking()
}

func (b *Bot) StartBlocking() {
	b.tb.Start()
}
