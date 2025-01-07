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
	ID         int64 `bson:"_id"`
	Username   string
	PlayerRole permissions.PlayerRole

	Language string

	SquadID   *string
	GuildID   *string
	SquadRole permissions.SquadRole

	FirstSeen time.Time
	LastSeen  time.Time

	// Basic info
	Castle     Castle
	PlayerName string

	Level        int
	CurrentExp   int
	NextLevelExp int

	Rank int

	// Stats
	Str int
	Dex int
	Vit int

	DetailedStats map[string]int

	ProfileUpdatedAt time.Time

	// School
	Schools          map[string]int // SchoolID -> Level
	SchoolsUpdatedAt time.Time

	// Balance
	PlayerBalance    int
	BankBalance      int
	BalanceUpdatedAt time.Time
}

func (p Player) Mention() string {
	if p.Username != "" {
		return "@" + p.Username
	}
	return fmt.Sprintf("<a href=\"%d\">%s</a>", p.ID, p.PlayerName)
}

func (p Player) Updated() bool {
	return p.ProfileUpdatedAt.After(time.Now().Add(-ProfileUpdateInterval))
}

// Constants
var UsernameRegex = regexp.MustCompile("^\\w{4,32}$")
var ProfileUpdateInterval = 48 * time.Hour

// Errors
var ErrPlayerNotFound = errors.New("player not found")
var ErrInvalidUsername = errors.New("invalid username")
var ErrInvalidText = errors.New("invalid text")
var ErrNeedProfileUpdate = errors.New("need profile update")
var ErrLeaderCannotLeave = errors.New("leader cannot leave squad")

// Interfaces
type PlayerUsecase interface {
	Create(ctx context.Context, scope permissions.Scope, id int64, username string) (*Player, error)

	Get(ctx context.Context, scope permissions.Scope, id int64) (*Player, error)
	GetByUsername(ctx context.Context, scope permissions.Scope, username string) (*Player, error)

	ParseMe(ctx context.Context, scope permissions.Scope, me string) (*Player, error)
	ParseHero(ctx context.Context, scope permissions.Scope, hero string) (*Player, error)
	ParseSchool(ctx context.Context, scope permissions.Scope, school string) (*Player, error)

	SquadAdd(ctx context.Context, scope permissions.Scope, id int64) (*Player, error)
	SquadRemove(ctx context.Context, scope permissions.Scope, id int64) (*Player, error)

	Seen(ctx context.Context, scope permissions.Scope, username string) (*Player, error)
}

type PlayerRepository interface {
	Create(ctx context.Context, obj *Player) error

	Get(ctx context.Context, id int64) (*Player, error)
	GetByUsername(ctx context.Context, username string) (*Player, error)

	ListBySquad(ctx context.Context, squadID string) ([]*Player, error)
	ListByGuild(ctx context.Context, guildID string) ([]*Player, error)

	CountBySquad(ctx context.Context, squadID string) (int64, error)
	CountByGuild(ctx context.Context, guildID string) (int64, error)

	Update(ctx context.Context, obj *Player) error

	Delete(ctx context.Context, id int64) error
}
