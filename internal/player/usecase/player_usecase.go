package usecase

import (
	"github.com/Russia9/Muskrat/pkg/domain"
)

type uc struct {
	repo domain.PlayerRepository
}

func NewPlayerUsecase(repo domain.PlayerRepository) domain.PlayerUsecase {
	return &uc{repo}
}
