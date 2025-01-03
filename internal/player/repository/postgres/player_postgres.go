package postgres

import (
	"context"

	"github.com/Russia9/Muskrat/internal/sqlc"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type repo struct {
	queries *sqlc.Queries
}

func sqlcToDomain(in sqlc.Player) *domain.Player {
	obj := domain.Player{
		ID:         in.ID,
		Username:   in.Username,
		PlayerRole: permissions.PlayerRole(in.PlayerRole),

		Language: in.Language,

		SquadID:   nil,
		GuildID:   nil,
		SquadRole: permissions.SquadRole(in.SquadRole),

		FirstSeen: in.FirstSeen.Time,
		LastSeen:  in.LastSeen.Time,

		Castle:     domain.Castle(in.Castle),
		PlayerName: in.PlayerName,

		Level:        int(in.Level),
		CurrentExp:   int(in.CurrentExp),
		NextLevelExp: int(in.NextLevelExp),

		Rank: int(in.Rank),

		Str: int(in.Str),
		Dex: int(in.Dex),
		Vit: int(in.Vit),

		DetailedStats: in.DetailedStats,

		ProfileUpdatedAt: in.ProfileUpdatedAt.Time,

		Schools:          in.Schools,
		SchoolsUpdatedAt: in.SchoolsUpdatedAt.Time,

		PlayerBalance:    int(in.PlayerBalance),
		BankBalance:      int(in.BankBalance),
		BalanceUpdatedAt: in.BalanceUpdatedAt.Time,
	}

	if in.SquadID.Valid {
		uuid := in.SquadID.String()
		obj.SquadID = &uuid
	}
	if in.GuildID.Valid {
		uuid := in.GuildID.String()
		obj.GuildID = &uuid
	}

	return &obj
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
		FirstSeen:  pgtype.Timestamptz{Time: obj.FirstSeen, Valid: true},
		LastSeen:   pgtype.Timestamptz{Time: obj.LastSeen, Valid: true},
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
	obj, err := r.queries.GetPlayer(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrPlayerNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "sqlc")
	}

	return sqlcToDomain(obj), nil
}

func (r *repo) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	obj, err := r.queries.GetPlayerByUsername(ctx, username)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrPlayerNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "sqlc")
	}

	return sqlcToDomain(obj), nil
}

func (r *repo) Update(ctx context.Context, obj *domain.Player) error {
	err := r.queries.UpdatePlayer(ctx, sqlc.UpdatePlayerParams{
		ID:         obj.ID,
		Username:   obj.Username,
		PlayerRole: int32(obj.PlayerRole),

		Language: obj.Language,

		SquadID:   pgtype.UUID{},
		GuildID:   pgtype.UUID{},
		SquadRole: sqlc.SquadRole(obj.SquadRole),

		FirstSeen: pgtype.Timestamptz{Time: obj.FirstSeen, Valid: true},
		LastSeen:  pgtype.Timestamptz{Time: obj.LastSeen, Valid: true},

		Castle:     sqlc.Castle(obj.Castle),
		PlayerName: obj.PlayerName,

		Level:        int32(obj.Level),
		CurrentExp:   int32(obj.CurrentExp),
		NextLevelExp: int32(obj.NextLevelExp),

		Rank: int32(obj.Rank),

		Str: int32(obj.Str),
		Dex: int32(obj.Dex),
		Vit: int32(obj.Vit),

		DetailedStats: obj.DetailedStats,

		ProfileUpdatedAt: pgtype.Timestamptz{Time: obj.ProfileUpdatedAt, Valid: true},

		Schools:          obj.Schools,
		SchoolsUpdatedAt: pgtype.Timestamptz{Time: obj.SchoolsUpdatedAt, Valid: true},

		PlayerBalance:    int32(obj.PlayerBalance),
		BankBalance:      int32(obj.BankBalance),
		BalanceUpdatedAt: pgtype.Timestamptz{Time: obj.BalanceUpdatedAt, Valid: true},
	})

	if err != nil {
		return errors.Wrap(err, "sqlc")
	}

	return nil
}
