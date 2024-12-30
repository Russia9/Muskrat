package usecase_test

import (
	"testing"

	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/google/go-cmp/cmp"
)

func assert(t *testing.T, expected *domain.Player, got *domain.Player) {
	if expected == nil && got == nil {
		return
	}
	if expected == nil || got == nil {
		t.Errorf("expected %v, got %v", expected, got)
		return
	}

	if !cmp.Equal(got, expected) {
		t.Error(cmp.Diff(got, expected))
	}
}
