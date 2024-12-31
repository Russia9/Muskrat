package postgres

import (
	"context"
	"encoding/json"

	"github.com/Russia9/Muskrat/internal/sqlc"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type repo struct {
	queries *sqlc.Queries
}

func sqlcToDomain(sqlcPlayer sqlc.Player) (*domain.Player, error) {
	obj := domain.Player{
		ID:         sqlcPlayer.ID,
		Username:   sqlcPlayer.Username,
		PlayerRole: permissions.PlayerRole(sqlcPlayer.PlayerRole),

		Language: sqlcPlayer.Language,

		SquadID:   nil,
		GuildID:   nil,
		SquadRole: permissions.SquadRole(sqlcPlayer.SquadRole),

		FirstSeen: sqlcPlayer.FirstSeen.Time,
		LastSeen:  sqlcPlayer.LastSeen.Time,

		Castle:     domain.Castle(sqlcPlayer.Castle),
		PlayerName: sqlcPlayer.PlayerName,

		Level:        int(sqlcPlayer.Level),
		CurrentExp:   int(sqlcPlayer.CurrentExp),
		NextLevelExp: int(sqlcPlayer.NextLevelExp),

		Rank: int(sqlcPlayer.Rank),

		Str: int(sqlcPlayer.Str),
		Dex: int(sqlcPlayer.Dex),
		Vit: int(sqlcPlayer.Vit),

		DetailedStats: nil,

		ProfileUpdatedAt: sqlcPlayer.ProfileUpdatedAt.Time,

		Schools:          nil,
		SchoolsUpdatedAt: sqlcPlayer.SchoolsUpdatedAt.Time,

		PlayerBalance:    int(sqlcPlayer.PlayerBalance),
		BankBalance:      int(sqlcPlayer.BankBalance),
		BalanceUpdatedAt: sqlcPlayer.BalanceUpdatedAt.Time,
	}

	if sqlcPlayer.SquadID.Valid {
		uuid := sqlcPlayer.SquadID.String()
		obj.SquadID = &uuid
	}
	if sqlcPlayer.GuildID.Valid {
		uuid := sqlcPlayer.GuildID.String()
		obj.GuildID = &uuid
	}

	var err error
	if sqlcPlayer.DetailedStats != nil {
		err = json.Unmarshal(sqlcPlayer.DetailedStats, &obj.DetailedStats)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal")
		}
	}

	if sqlcPlayer.Schools != nil {
		err = json.Unmarshal(sqlcPlayer.Schools, &obj.Schools)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal")
		}
	}

	return &obj, nil
}

func NewPlayerRepository(db *pgxpool.Pool) domain.PlayerRepository {
	return &repo{sqlc.New(db)}
}

func (r *repo) Create(ctx context.Context, obj *domain.Player) error {
	err := r.queries.CreatePlayer(ctx, sqlc.CreatePlayerParams{
		ID:         obj.ID,
		Username:   obj.Username,
		PlayerRole: int32(obj.PlayerRole),
		Language:   obj.Language,
		FirstSeen:  pgtype.Timestamptz{Time: obj.FirstSeen},
		LastSeen:   pgtype.Timestamptz{Time: obj.LastSeen},
	})
	if err != nil {
		return errors.Wrap(err, "sqlc")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeletePlayer(ctx, id)
	if err != nil {
		return errors.Wrap(err, "sqlc")
	}

	return nil
}

func (r *repo) Get(ctx context.Context, id int64) (*domain.Player, error) {
	panic("unimplemented")
}

func (r *repo) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	panic("unimplemented")
}

func (r *repo) Update(ctx context.Context, obj *domain.Player) error {
	panic("unimplemented")
}
