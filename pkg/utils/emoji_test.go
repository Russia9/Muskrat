package utils_test

import (
	"testing"

	"github.com/Russia9/Muskrat/pkg/utils"
)

func TestDigitToKeycap(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want string
	}{
		{
			name: "0",
			x:    0,
			want: "0️⃣",
		},
		{
			name: "1",
			x:    1,
			want: "1️⃣",
		},
		{
			name: "2",
			x:    2,
			want: "2️⃣",
		},
		{
			name: "9",
			x:    9,
			want: "9️⃣",
		},
		{
			name: "invalid",
			x:    10,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.DigitToKeycap(tt.x); got != tt.want {
				t.Errorf("DigitToKeycap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeycapToDigit(t *testing.T) {
	tests := []struct {
		name string
		x    string
		want int
	}{
		{
			name: "0",
			x:    "0️⃣",
			want: 0,
		},
		{
			name: "1",
			x:    "1️⃣",
			want: 1,
		},
		{
			name: "2",
			x:    "2️⃣",
			want: 2,
		},
		{
			name: "9",
			x:    "9️⃣",
			want: 9,
		},
		{
			name: "invalid",
			x:    "a",
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.KeycapToDigit(tt.x); got != tt.want {
				t.Errorf("KeycapToDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
