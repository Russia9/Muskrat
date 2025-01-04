package squad

import (
	"context"
	"regexp"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Module struct {
	player domain.PlayerUsecase
	squad  domain.SquadUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase, squad domain.SquadUsecase) *Module {
	m := &Module{player, squad, tb, l}

	tb.Handle("/squad_create", m.SquadCreate)

	return m
}

var SquadCreateRegex = regexp.MustCompile("/squad_create (.{1,32})")

func (m *Module) SquadCreate(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check if chat is supergroup
	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Send(m.l.Text(c, "squad_create_not_in_chat"))
	}

	// Check if caller is a chat admin
	cm, err := m.tb.ChatMemberOf(c.Chat(), c.Sender())
	if err != nil {
		return errors.Wrap(err, "telebot")
	}
	if cm.Role != telebot.Administrator && cm.Role != telebot.Creator {
		return c.Send(m.l.Text(c, "squad_create_not_chat_admin"))
	}

	// Check if message is in correct format
	if !SquadCreateRegex.MatchString(c.Text()) {
		return c.Send(m.l.Text(c, "squad_create_wrong_format"))
	}

	// Check if chat already has a squad
	_, err = m.squad.GetByChatID(context.Background(), permissions.InternalScope, c.Chat().ID)
	if err == nil {
		return c.Send(m.l.Text(c, "squad_chat_already_attached"))
	} else if !errors.Is(err, domain.ErrSquadNotFound) {
		return errors.Wrap(err, "squad uc")
	}

	sq, err := m.squad.Create(context.Background(), scope, c.Chat().ID, SquadCreateRegex.FindStringSubmatch(c.Text())[1])
	if errors.Is(err, domain.ErrAlreadyInSquad) {
		return c.Send(m.l.Text(c, "squad_create_already_in_squad"))
	} else if errors.Is(err, domain.ErrNeedProfileUpdate) {
		return c.Send(m.l.Text(c, "need_profile_update"))
	} else if err != nil {
		return errors.Wrap(err, "squad uc")
	}

	return c.Send(m.l.Text(c, "squad_create_success", squadCreateTmpl{sq.Name, sq.ChatID}))
}

type squadCreateTmpl struct {
	Name   string
	ChatID int64
}
