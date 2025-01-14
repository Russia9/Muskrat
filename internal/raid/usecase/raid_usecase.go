package usecase

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/pkg/errors"
	"time"
)

type uc struct {
	repo domain.RaidRepository
}

func NewRaidUsecase(repository domain.RaidRepository) domain.RaidUsecase {
	return &uc{repo: repository}
}

func (u uc) UpdateOrCreate(ctx context.Context, name string, cell string, preRaidTime int32) error {
	raidTime := time.Now().UTC().Add(time.Duration(preRaidTime) * time.Minute)
	raidInfo := &domain.Raid{
		Name: name,
		Cell: cell,
		Time: raidTime,
	}
	err := u.repo.UpdateOrCreate(ctx, raidInfo)
	if err != nil {
		return errors.Wrap(err, "update or create raid ")
	}
	return nil
}

func (u uc) List(ctx context.Context) ([]*domain.Raid, error) {
	raids, err := u.repo.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list raid ")
	}
	return raids, nil
}
