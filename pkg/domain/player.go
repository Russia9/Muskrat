package domain

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/Russia9/Muskrat/pkg/permissions"
)

// Entity
type Player struct {
	ID         int64                  `json:"id"`
	Username   string                 `json:"username"`
	PlayerRole permissions.PlayerRole `json:"player_role"`

	Language string `json:"language"`

	SquadID   *string               `json:"squad_id"`
	GuildID   *string               `json:"guild_id"`
	SquadRole permissions.SquadRole `json:"squad_role"`

	FirstSeen time.Time `json:"first_seen"`
	LastSeen  time.Time `json:"last_seen"`

	// Basic info
	Castle     Castle `json:"castle"`
	PlayerName string `json:"player_name"`

	Level        int `json:"level"`
	CurrentExp   int `json:"current_exp"`
	NextLevelExp int `json:"next_level_exp"`

	Rank int `json:"rank"`

	// Stats
	Str int `json:"str"`
	Dex int `json:"dex"`
	Vit int `json:"vit"`

	DetailedStats map[string]int `json:"detailed_stats"`

	ProfileUpdatedAt time.Time `json:"profile_updated_at"`

	// School
	Schools          map[string]int `json:"schools"` // SchoolID -> Level
	SchoolsUpdatedAt time.Time      `json:"schools_updated_at"`

	// Balance
	PlayerBalance    int       `json:"player_balance"`
	BankBalance      int       `json:"bank_balance"`
	BalanceUpdatedAt time.Time `json:"balance_updated_at"`
}

func (p Player) Mention() string {
	if p.Username != "" {
		return "@" + p.Username
	}
	return fmt.Sprintf("<a href=\"%d\">%s</a>", p.ID, p.PlayerName)
}

// Constants
var UsernameRegex = regexp.MustCompile("^\\w{4,32}$")

// Errors
var ErrPlayerNotFound = errors.New("player not found")
var ErrInvalidUsername = errors.New("invalid username")
var ErrInvalidText = errors.New("invalid text")

// Interfaces
type PlayerUsecase interface {
	Create(ctx context.Context, scope permissions.Scope, id int64, username string) (*Player, error)

	Get(ctx context.Context, scope permissions.Scope, id int64) (*Player, error)
	GetByUsername(ctx context.Context, scope permissions.Scope, username string) (*Player, error)

	ParseMe(ctx context.Context, scope permissions.Scope, me string) (*Player, error)
	ParseHero(ctx context.Context, scope permissions.Scope, hero string) (*Player, error)
	ParseSchool(ctx context.Context, scope permissions.Scope, school string) (*Player, error)

	Seen(ctx context.Context, scope permissions.Scope, username string) (*Player, error)
}

type PlayerRepository interface {
	Create(ctx context.Context, obj *Player) error

	Get(ctx context.Context, id int64) (*Player, error)
	GetByUsername(ctx context.Context, username string) (*Player, error)

	Update(ctx context.Context, obj *Player) error

	Delete(ctx context.Context, id int64) error
}
