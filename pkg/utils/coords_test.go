package utils_test

import (
	"testing"

	"github.com/Russia9/Muskrat/pkg/utils"
)

func TestCoordsToGoto(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"G 3#3", "g_3_3"},
		{"GY 5", "gy_5"},
		{"0#0", "0_0"},
	}
	for _, test := range tests {
		actual := utils.CoordsToGoto(test.input)
		if actual != test.expected {
			t.Errorf("CoordsToGoto(%q) = %q, want %q", test.input, actual, test.expected)
		}
	}
}

func TestGotoToCoords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"G 3#3", "g_3_3"},
		{"GY 5", "gy_5"},
		{"0#0", "0_0"},
	}
	for _, test := range tests {
		actual := utils.GotoToCoords(test.input)
		if actual != test.expected {
			t.Errorf("GotoToCoords(%q) = %q, want %q", test.input, actual, test.expected)
		}
	}
}
