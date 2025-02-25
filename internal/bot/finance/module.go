package finance

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/Russia9/Muskrat/pkg/utils"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Module struct {
	player domain.PlayerUsecase
	squad  domain.SquadUsecase
	guild  domain.GuildUsecase

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase, squad domain.SquadUsecase, guild domain.GuildUsecase) *Module {
	m := &Module{player, squad, guild, tb, l}

	tb.Handle("💰 Finance", m.Finance)
	tb.Handle("💰 Финансы", m.Finance)
	tb.Handle("/finance", m.Finance)
	tb.Handle("\ffinance", m.Finance)

	return m
}

func (m *Module) Finance(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Determine finance type
	financeType := "squad"
	if c.Data() == "guild" {
		financeType = "guild"
	}

	// Check if player is in squad
	if scope.SquadRole < permissions.SquadRoleMember || scope.SquadID == nil {
		return c.Send(m.l.Text(c, "squad_not_in_squad"))
	}

	if financeType == "guild" {
		// Check if player is in guild
		if scope.GuildRole < permissions.SquadRoleMember || scope.GuildID == nil {
			return c.Send(m.l.Text(c, "guild_not_in_guild"))
		}
	}

	// Get squad or guild name
	name := ""
	if financeType == "guild" {
		// Get Guild
		g, err := m.guild.Get(context.Background(), scope, *scope.GuildID)
		if err != nil {
			return err
		}
		name = g.Name
	} else {
		// Get Squad
		s, err := m.squad.Get(context.Background(), scope, *scope.SquadID)
		if err != nil {
			return err
		}
		name = s.Name
	}

	// Get player list
	var players []*domain.Player
	var err error
	if financeType == "guild" {
		players, err = m.player.ListByGuild(context.Background(), scope, *scope.GuildID, domain.PlayerSortBalance)
	} else {
		players, err = m.player.ListBySquad(context.Background(), scope, *scope.SquadID, domain.PlayerSortBalance)
	}
	if err != nil {
		return err
	}

	// Generate template
	t := template{
		Name:    name,
		Players: make([]playerFinance, len(players)),
	}

	namePads := 0
	playerBalancePads := 0
	bankBalancePads := 0

	for i, p := range players {
		// Add player to template
		t.Players[i] = playerFinance{
			Name:          "",
			PlayerBalance: p.PlayerBalance,
			BankBalance:   p.BankBalance,
		}

		// Determine player name (Username or PlayerName)
		if p.Username != "" {
			t.Players[i].Name = "@" + utils.ShortString(p.Username, 12)
		} else {
			t.Players[i].Name = utils.ShortString(p.PlayerName, 12)
		}

		// Update padding
		if len(t.Players[i].Name) > namePads {
			namePads = len(t.Players[i].Name)
		}
		if int(math.Ceil(math.Log10(float64(p.PlayerBalance)))) > playerBalancePads {
			playerBalancePads = int(math.Ceil(math.Log10(float64(p.PlayerBalance))))
		}
		if int(math.Ceil(math.Log10(float64(p.BankBalance)))) > bankBalancePads {
			bankBalancePads = int(math.Ceil(math.Log10(float64(p.BankBalance))))
		}

		// Add balance to total and count last update time
		t.TotalBalance += p.PlayerBalance + p.BankBalance
		if p.BalanceUpdatedAt.After(time.Now().Add(-time.Hour * 24)) {
			t.CurrentBalance += p.PlayerBalance + p.BankBalance
			t.Players[i].Time = utils.FormatDuration(time.Since(p.BalanceUpdatedAt))
		} else {
			if p.BalanceUpdatedAt.IsZero() {
				t.Players[i].Time = "❗️N/A"
			} else {
				t.Players[i].Time = "❗️" + utils.FormatDuration(time.Since(p.BalanceUpdatedAt))
			}
		}
	}

	// Set formatting for padding
	t.NameFmt = fmt.Sprintf("%%-%ds", namePads)
	t.PlayerBalanceFmt = fmt.Sprintf("%%%dd", playerBalancePads)
	t.BankBalanceFmt = fmt.Sprintf("%%%dd", bankBalancePads)

	// Generate markup
	markup := &telebot.ReplyMarkup{}
	if c.Chat().Type == telebot.ChatPrivate && scope.GuildRole > permissions.SquadRoleNone && scope.GuildID != nil {
		markup.Inline(markup.Row(
			markup.Data(m.l.Text(c, "finance_squad_btn"), "finance", "squad"),
			markup.Data(m.l.Text(c, "finance_guild_btn"), "finance", "guild"),
		))
	}

	return c.EditOrSend(m.l.Text(c, "finance_"+financeType, t), markup)
}

type playerFinance struct {
	Name string

	PlayerBalance int
	BankBalance   int

	Time string
}

type template struct {
	Name    string
	Players []playerFinance

	NameFmt          string
	PlayerBalanceFmt string
	BankBalanceFmt   string

	TotalBalance   int
	CurrentBalance int // Last update 48h or less
}
