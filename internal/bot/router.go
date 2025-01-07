package bot

import (
	"github.com/Russia9/Muskrat/pkg/utils"
	"gopkg.in/telebot.v3"
)

func (b *Bot) Router(c telebot.Context) error {
	// Get Player from context
	// player := c.Get("player").(*domain.Player)

	if c.Message().IsForwarded() && c.Message().OriginalSender != nil && c.Message().OriginalSender.ID == utils.ChatWarsBot {
		return b.parse.Router(c)
	}

	return b.Menu(c)
}
