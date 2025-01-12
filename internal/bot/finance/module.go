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

	tb *telebot.Bot
	l  *layout.Layout
}

func NewModule(tb *telebot.Bot, l *layout.Layout, player domain.PlayerUsecase, squad domain.SquadUsecase) *Module {
	m := &Module{player, squad, tb, l}

	tb.Handle("💰 Finance", m.Finance)
	tb.Handle("💰 Финансы", m.Finance)
	tb.Handle("/finance_squad", m.Finance)

	return m
}

func (m *Module) Finance(c telebot.Context) error {
	scope := c.Get("scope").(permissions.Scope)

	// Check if player is in squad
	if scope.SquadRole < permissions.SquadRoleMember || scope.SquadID == nil {
		return c.Send(m.l.Text(c, "squad_not_in_squad"))
	}

	// Get player list
	players, err := m.player.ListBySquad(context.Background(), scope, *scope.SquadID)
	if err != nil {
		return err
	}

	// Generate template
	t := template{
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
			t.Players[i].Name = "@" + p.Username
		} else {
			t.Players[i].Name = p.PlayerName
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
		if p.BalanceUpdatedAt.After(time.Now().Add(-time.Hour * 48)) {
			t.CurrentBalance += p.PlayerBalance + p.BankBalance
			t.Players[i].Time = utils.FormatDuration(time.Now().Sub(p.BalanceUpdatedAt))
		} else {
			if p.BalanceUpdatedAt.IsZero() {
				t.Players[i].Time = "❗️N/A"
			} else {
				t.Players[i].Time = "❗️" + utils.FormatDuration(time.Now().Sub(p.BalanceUpdatedAt))
			}
		}
	}

	// Set formatting for padding
	t.NameFmt = fmt.Sprintf("%%-%ds", namePads)
	t.PlayerBalanceFmt = fmt.Sprintf("%%%dd", playerBalancePads)
	t.BankBalanceFmt = fmt.Sprintf("%%%dd", bankBalancePads)

	return c.Send(m.l.Text(c, "finance_squad", t))
}

type playerFinance struct {
	Name string

	PlayerBalance int
	BankBalance   int

	Time string
}

type template struct {
	Players []playerFinance

	NameFmt          string
	PlayerBalanceFmt string
	BankBalanceFmt   string

	TotalBalance   int
	CurrentBalance int // Last update 48h or less
}
