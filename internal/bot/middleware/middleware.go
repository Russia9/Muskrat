package middleware

import (
	"github.com/Russia9/Muskrat/pkg/domain"
	"gopkg.in/telebot.v3/layout"
)

type Middleware struct {
	player domain.PlayerUsecase

	layout *layout.Layout
}

func NewMiddleware(player domain.PlayerUsecase, layout *layout.Layout) *Middleware {
	return &Middleware{player, layout}
}
