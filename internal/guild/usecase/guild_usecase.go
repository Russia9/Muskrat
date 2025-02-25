package usecase

import (
	"github.com/Russia9/Muskrat/pkg/domain"
)

type uc struct {
	repo   domain.GuildRepository
	player domain.PlayerRepository
	squad  domain.SquadRepository
}

func NewGuildUsecase(repo domain.GuildRepository, player domain.PlayerRepository, squad domain.SquadRepository) domain.GuildUsecase {
	return &uc{repo, player, squad}
}
