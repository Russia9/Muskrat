package usecase

import (
	"context"
	"regexp"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/Russia9/Muskrat/pkg/permissions"
)

var schoolRegex = regexp.MustCompile(" (...) .*\n /school_(.*)")

func (u *uc) ParseSchool(ctx context.Context, scope permissions.Scope, school string) (*domain.Player, error) {
	panic("unimplemented")
}
