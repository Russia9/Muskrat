package bot

import (
	"gopkg.in/telebot.v3"
)

func (b bot) Router(c telebot.Context) error {
	// Get Player from context
	// player := c.Get("player").(*domain.Player)

	return b.Menu(c)
}
