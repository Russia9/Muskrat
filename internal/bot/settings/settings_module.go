package settings

import (
	"context"
	"errors"
	"regexp"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Module struct {
	player domain.PlayerUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase) *Module {
	m := &Module{player, tb, l}

	tb.Handle("/lang", m.HandleLocale)

	return m
}

var LangRegex = regexp.MustCompile(`^/lang (en|ru)$`)

func (m *Module) HandleLocale(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check input
	if !LangRegex.MatchString(c.Text()) {
		return c.Send(m.l.Text(c, "settings_locale_wrong_format"))
	}

	// Parse command
	lang := LangRegex.FindStringSubmatch(c.Text())[1]

	// Change language
	pl, err := m.player.Locale(context.Background(), scope, lang)
	if errors.Is(err, domain.ErrUnsupportedLanguage) {
		return c.Send(m.l.Text(c, "settings_locale_wrong_format"))
	} else if err != nil {
		return err
	}
	m.l.SetLocale(c, pl.Locale)

	return c.Send(m.l.Text(c, "settings_locale_success", pl))
}
