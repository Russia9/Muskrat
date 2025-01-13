package usecase

import (
	"context"
	"github.com/Russia9/Muskrat/pkg/domain"
	"time"
)

type uc struct {
	repo domain.RaidRepository
}

func NewRaidUsecase(repository domain.RaidRepository) domain.RaidUsecase {
	return &uc{repo: repository}
}

func (u uc) Create(ctx context.Context, name string, cell string, preRaidTime int32) error {
	raidTime := time.Now().UTC().Add(time.Duration(preRaidTime) * time.Minute)
	raidInfo := &domain.Raid{Name: name, Cell: cell, Time: raidTime}
	return u.repo.Create(ctx, raidInfo)
}

func (u uc) Get(ctx context.Context, raid *domain.Raid) (*domain.Raid, error) {
	//TODO implement me
	panic("implement me")
}
