package squad

import (
	"context"
	"errors"
	"regexp"
	"strconv"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"gopkg.in/telebot.v3"
)

var SquadKickRegex = regexp.MustCompile(`/squad_kick(?:@MuskratBot)?(?:\s(?:@(.*)|(\d+)))?`)

func (m *Module) SquadKick(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check if message is in correct format
	if !SquadKickRegex.MatchString(c.Text()) {
		return c.Send(m.l.Text(c, "squad_kick_wrong_format"))
	}

	// Parse target
	target := SquadKickRegex.FindStringSubmatch(c.Text())

	// Get player based on target
	var err error
	var pl *domain.Player
	switch {
	case target[1] != "":
		pl, err = m.player.GetByUsername(context.Background(), scope, target[1])
	case target[2] != "":
		var id int64
		id, err = strconv.ParseInt(target[2], 10, 64)
		if err != nil {
			return c.Send(m.l.Text(c, "squad_kick_wrong_format"))
		}

		pl, err = m.player.Get(context.Background(), scope, id)
	case c.Message().ReplyTo != nil && c.Message().ReplyTo.Sender != nil:
		pl, err = m.player.Get(context.Background(), scope, c.Message().ReplyTo.Sender.ID)
	default:
		return c.Send(m.l.Text(c, "squad_kick_wrong_format"))
	}
	if err != nil {
		return err
	}

	// Remove player from squad
	pl, err = m.player.SquadRemove(context.Background(), scope, pl.ID)
	if errors.Is(err, domain.ErrNotInSquad) {
		return c.Send(m.l.Text(c, "squad_kick_not_in_squad"))
	} else if errors.Is(err, domain.ErrLeaderCannotLeave) {
		return c.Send(m.l.Text(c, "squad_kick_leader_can_not_leave"))
	} else if err != nil {
		return err
	}

	return c.Send(m.l.Text(c, "squad_kick_success", pl))
}
