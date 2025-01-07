package squad

import "gopkg.in/telebot.v3"

func (m *Module) SquadHelp(c telebot.Context) error {
	return c.Send(m.l.Text(c, "squad_help"))
}
