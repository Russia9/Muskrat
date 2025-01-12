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

	Locale string

	SquadID   *string
	SquadRole permissions.SquadRole

	GuildID   *string
	GuildRole permissions.SquadRole

	FirstSeen time.Time
	LastSeen  time.Time

	// Basic info
	Castle     Castle
	PlayerName string

	Level        int
	CurrentExp   int
	NextLevelExp int

	BasicsUpdatedAt time.Time

	// Stats
	Rank int

	Str int
	Dex int
	Vit int

	DetailedStats map[string]int

	StatsUpdatedAt time.Time

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
	if p.PlayerName != "" {
		return fmt.Sprintf("<a href=\"%d\">%s</a>", p.ID, p.PlayerName)
	}
	return fmt.Sprintf("<a href=\"%d\">#%d</a>", p.ID, p.ID)
}

func (p Player) Updated() bool {
	return p.BasicsUpdatedAt.After(time.Now().Add(-ProfileUpdateInterval))
}

// Constants
var UsernameRegex = regexp.MustCompile("^\\w{4,32}$")
var ProfileUpdateInterval = 48 * time.Hour

const DefaultLocale = "ru"

var Locales = []string{"en", "ru"}

type PlayerSort int

const (
	PlayerSortLevel PlayerSort = iota
	PlayerSortRank
	PlayerSortBalance
)

// Errors
var ErrPlayerNotFound = errors.New("player not found")
var ErrInvalidUsername = errors.New("invalid username")
var ErrInvalidText = errors.New("invalid text")
var ErrNeedProfileUpdate = errors.New("need profile update")
var ErrLeaderCannotLeave = errors.New("leader cannot leave squad")
var ErrUnsupportedLanguage = errors.New("unsupported language")

// Interfaces
type PlayerUsecase interface {
	Create(ctx context.Context, scope permissions.Scope, id int64, username string) (*Player, error)

	Get(ctx context.Context, scope permissions.Scope, id int64) (*Player, error)
	GetByUsername(ctx context.Context, scope permissions.Scope, username string) (*Player, error)

	ListBySquad(ctx context.Context, scope permissions.Scope, squadID string, sort PlayerSort) ([]*Player, error)

	ParseMe(ctx context.Context, scope permissions.Scope, me string) (*Player, error)
	ParseHero(ctx context.Context, scope permissions.Scope, hero string) (*Player, error)
	ParseSchool(ctx context.Context, scope permissions.Scope, school string) (*Player, error)

	SquadAdd(ctx context.Context, scope permissions.Scope, id int64) (*Player, error)
	SquadRemove(ctx context.Context, scope permissions.Scope, id int64) (*Player, error)

	Locale(ctx context.Context, scope permissions.Scope, locale string) (*Player, error)

	Seen(ctx context.Context, scope permissions.Scope, username string) (*Player, error)
}

type PlayerRepository interface {
	Create(ctx context.Context, obj *Player) error

	Get(ctx context.Context, id int64) (*Player, error)
	GetByUsername(ctx context.Context, username string) (*Player, error)

	ListBySquad(ctx context.Context, squadID string, sort PlayerSort) ([]*Player, error)
	ListByGuild(ctx context.Context, guildID string, sort PlayerSort) ([]*Player, error)

	CountBySquad(ctx context.Context, squadID string) (int64, error)
	CountByGuild(ctx context.Context, guildID string) (int64, error)

	Update(ctx context.Context, obj *Player) error

	Delete(ctx context.Context, id int64) error
}
