package parse

import (
	"context"
	"errors"
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
	guild  domain.GuildUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase, guild domain.GuildUsecase) *Module {
	return &Module{player, guild, tb, l}
}

var meRegex = regexp.MustCompile(`^Region Battle in [\w\W]*`)
var heroRegex = regexp.MustCompile(`[\w\W]*⚙️Settings /settings[\w\W]*`)
var schoolRegex = regexp.MustCompile(`^📚School Management[\w\W]*`)
var idlistRegex = regexp.MustCompile(`👣\d+ (\d+)`)
var guildRegex = regexp.MustCompile(`[🇮🇲🇻🇦🇪🇺🇲🇴]+(?:\[(.+)\] )?([\w ]*)(?:\n📍Guild HQ: .*\[(.*)\])?[\w\W]+🏅Level: (\d+)`)

var ErrNotForwarded = errors.New("not forwarded")

const TimeTreshold = 40 * time.Second

func (m *Module) Router(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	if !c.Message().IsForwarded() || c.Message().OriginalSender.ID != utils.ChatWarsBot {
		return ErrNotForwarded
	}

	// Check message time
	ogTime := time.Unix(int64(c.Message().OriginalUnixtime), 0)
	if time.Since(ogTime) > TimeTreshold { // If message is older than TimeTreshold
		if c.Chat().Type == telebot.ChatPrivate {
			return c.Reply(m.l.Text(c, "parse_too_old"))
		}

		return nil
	}

	// Parse message if possible
	var err error
	switch {
	case meRegex.MatchString(c.Text()):
		_, err = m.player.ParseMe(context.Background(), scope, c.Text())
		if err != nil {
			return err
		}

		return m.react(c)
	case heroRegex.MatchString(c.Text()):
		_, err = m.player.ParseHero(context.Background(), scope, c.Text())
		if err != nil {
			return err
		}

		return m.react(c)
	case schoolRegex.MatchString(c.Text()):
		_, err = m.player.ParseSchool(context.Background(), scope, c.Text())
		if err != nil {
			return err
		}

		return m.react(c)
	case idlistRegex.MatchString(c.Text()):
		if scope.GuildRole == permissions.SquadRoleLeader && scope.GuildID != nil {
			_, err = m.guild.ParseList(context.Background(), scope, c.Text())
			if err != nil {
				return err
			}

			return m.react(c)
		}
	case guildRegex.MatchString(c.Text()):
		if scope.GuildRole == permissions.SquadRoleLeader && scope.GuildID != nil {
			_, err = m.guild.ParseGuild(context.Background(), scope, c.Text())
			if err != nil {
				return err
			}

			return m.react(c)
		}
	}

	return nil
}

const PreferredEmoji = "✍️"

func (m *Module) react(c telebot.Context) error {
	// Check if chat is PM
	if c.Chat().Type == telebot.ChatPrivate {
		// Delete message if PM
		return c.Delete()
	}

	// Try to set the reaction for message
	_ = c.Bot().React(c.Chat(), c.Message(), telebot.ReactionOptions{
		Reactions: []telebot.Reaction{
			{
				Type:  "emoji",
				Emoji: PreferredEmoji,
			},
		},
	})

	return nil
}
