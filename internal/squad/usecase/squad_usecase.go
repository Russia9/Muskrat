package usecase

import (
	"github.com/Russia9/Muskrat/pkg/domain"
)

type uc struct {
	repo   domain.SquadRepository
	player domain.PlayerRepository
}

func NewSquadUsecase(repo domain.SquadRepository, player domain.PlayerRepository) domain.SquadUsecase {
	return &uc{repo, player}
}
