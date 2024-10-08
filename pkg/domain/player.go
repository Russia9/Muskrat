package domain

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Entity
type Player struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`

	SquadID   int64      `json:"squad_id"`
	GuildID   int64      `json:"guild_id"`
	SquadRole PlayerRole `json:"squad_role"`

	// Basic info
	PlayerName   string `json:"player_name"`
	Level        int    `json:"level"`
	CurrentExp   int    `json:"current_exp"`
	NextLevelExp int    `json:"next_level_exp"`

	Rank int `json:"rank"`

	// Stats
	Str int `json:"str"`
	Dex int `json:"dex"`
	Vit int `json:"vit"`

	AttackForce   int `json:"attack_force"`
	AttackSpeed   int `json:"attack_speed"`
	CriticalRate  int `json:"critical_rate"`
	CriticalForce int `json:"critical_force"`
	Accuracy      int `json:"accuracy"`
	Evasion       int `json:"evasion"`
	ArmorScore    int `json:"armor_score"`
	MoveSpeed     int `json:"move_speed"`

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
type PlayerRole int

const (
	PlayerRoleStranger PlayerRole = iota
	PlayerRoleMember
	PlayerRoleSquire
	PlayerRoleLeader
)

// Errors
var ErrPlayerNotFound = errors.New("player not found")

// Interfaces
type PlayerUsecase interface {
}

type PlayerRepository interface {
	Create(ctx context.Context, obj *Player) error

	Get(ctx context.Context, id int64) (*Player, error)
	GetByUsername(ctx context.Context, username string) (*Player, error)

	Update(ctx context.Context, obj *Player) error

	Delete(ctx context.Context, id int64) error
}
