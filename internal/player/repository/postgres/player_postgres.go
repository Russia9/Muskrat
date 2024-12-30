package postgres

import (
	"context"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewPlayerRepository(db *pgxpool.Pool) domain.PlayerRepository {
	return &repo{db}
}

func (r *repo) Create(ctx context.Context, obj *domain.Player) error {
	panic("unimplemented")
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	panic("unimplemented")
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
