package bot

import (
	"github.com/Russia9/Muskrat/internal/bot/middleware"
	"github.com/Russia9/Muskrat/internal/bot/parse"
	"github.com/Russia9/Muskrat/pkg/domain"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Bot struct {
	bot    *telebot.Bot
	layout *layout.Layout

	parse *parse.Module
}

func NewBot(tb *telebot.Bot, l *layout.Layout, pl domain.PlayerUsecase) *Bot {
	b := &Bot{
		bot:    tb,
		layout: l,
	}

	// Register middleware
	m := middleware.NewMiddleware(pl, l)
	b.bot.Use(middleware.Logger)
	b.bot.Use(m.Player)

	// Create Modules
	b.parse = parse.NewModule(tb, l, pl)

	// Register handlers
	b.bot.Handle(telebot.OnText, b.Router)

	return b
}

func (b *Bot) StartAsync() {
	go b.StartBlocking()
}

func (b *Bot) StartBlocking() {
	b.bot.Start()
}
