package squad

import (
	"context"
	"regexp"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
)

var SquadRenameRegex = regexp.MustCompile("/squad_rename(?:@MuskratBot)? (.{1,32})")

func (m *Module) SquadRename(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check if message is in correct format
	if !SquadRenameRegex.MatchString(c.Text()) {
		return c.Send(m.l.Text(c, "squad_rename_wrong_format"))
	}

	// Change squad name
	sq, err := m.squad.ChangeName(context.Background(), scope, SquadRenameRegex.FindStringSubmatch(c.Text())[1])
	if errors.Is(err, domain.ErrNotInSquad) {
		return c.Send(m.l.Text(c, "squad_not_in_squad"))
	} else if err != nil {
		return errors.Wrap(err, "squad uc")
	}

	return c.Send(m.l.Text(c, "squad_rename_success", sq))
}
