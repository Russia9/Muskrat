package bot

import (
	"gopkg.in/telebot.v3"
)

func (b *Bot) Menu(c telebot.Context) error {
	if c.Chat().Type != telebot.ChatPrivate {
		return nil
	}

	// Generate markup
	m := &telebot.ReplyMarkup{}
	m.Reply(
		m.Row(m.Text(b.l.Text(c, "menu_finance")), m.Text(b.l.Text(c, "menu_craft"))),
		m.Row(m.Text(b.l.Text(c, "menu_roster")), m.Text(b.l.Text(c, "menu_stock"))),
	)

	return c.Send(b.l.Text(c, "menu"), m)
}
