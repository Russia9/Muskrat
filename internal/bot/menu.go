package bot

import "gopkg.in/telebot.v3"

func (b bot) Menu(c telebot.Context) error {
	return c.Send("test")
}
