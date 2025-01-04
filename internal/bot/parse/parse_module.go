package parse

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/Russia9/Muskrat/pkg/utils"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Module struct {
	player domain.PlayerUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase) *Module {
	return &Module{player, tb, l}
}

var meRegex = regexp.MustCompile("^Region Battle in [\\w\\W]*")
var heroRegex = regexp.MustCompile("[\\w\\W]*⚙️Settings /settings[\\w\\W]*")
var schoolRegex = regexp.MustCompile("^📚School Management[\\w\\W]*")

var ErrNotForwarded = errors.New("not forwarded")

const TimeTreshold = 40 * time.Second

func (m *Module) Router(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	if !c.Message().IsForwarded() || c.Message().OriginalSender.ID != utils.ChatWarsBot {
		return ErrNotForwarded
	}

	// Check message time
	ogTime := time.Unix(int64(c.Message().OriginalUnixtime), 0)
	if time.Now().Sub(ogTime) > TimeTreshold {
		if c.Chat().Type == telebot.ChatPrivate {
			return c.Send(m.l.Text(c, "parse_too_old"))
		}
	}

	var err error
	switch {
	case meRegex.MatchString(c.Text()):
		_, err = m.player.ParseMe(context.Background(), scope, c.Text())
		if err != nil {
			return err
		}
	case heroRegex.MatchString(c.Text()):
		_, err = m.player.ParseHero(context.Background(), scope, c.Text())
		if err != nil {
			return err
		}
	case schoolRegex.MatchString(c.Text()):
		_, err = m.player.ParseSchool(context.Background(), scope, c.Text())
		if err != nil {
			return err
		}
	default:
		fmt.Println("1")
		return nil
	}

	return nil
}
