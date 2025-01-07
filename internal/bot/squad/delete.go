package squad

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/permissions"
	"gopkg.in/telebot.v3"
)

func (m *Module) SquadDelete(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	if scope.SquadRole != permissions.SquadRoleLeader || scope.SquadID == nil {
		return permissions.ErrForbidden
	}

	// Get Squad
	sq, err := m.squad.Get(context.Background(), scope, *scope.SquadID)
	if err != nil {
		return err
	}

	return c.Send(m.l.Text(c, "squad_delete_confirm", sq))
}

func (m *Module) SquadDeleteConfirm(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	err := m.squad.Delete(context.Background(), scope)
	if err != nil {
		return err
	}

	return c.Send(m.l.Text(c, "squad_delete_success"))
}
