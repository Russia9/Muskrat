package bot

import (
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type bot struct {
	bot    *telebot.Bot
	layout *layout.Layout
}
