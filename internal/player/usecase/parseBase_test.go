package usecase_test

import (
	"testing"
	"time"

	"github.com/Russia9/Muskrat/internal/player/usecase"
	"github.com/Russia9/Muskrat/pkg/domain"
	"github.com/agiledragon/gomonkey/v2"
)

func TestParseBase(t *testing.T) {
	// Mock time
	gomonkey.ApplyFuncReturn(time.Now, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC))

	tests := []struct {
		name    string
		msg     string
		want    domain.Player
		wantErr error
	}{
		{
			name: "Me",
			msg: `Region Battle in 8h 39 minutes!

🌟Поздравляем! Новый уровень!🌟
Нажми /level_up

🇲🇴[HC]M9co PakoB explorer of Green Castle
🏅Level: 30 6.24%
⚔️Rank: 1125
🛡️Armor: 131
📖Exp: 47154/54934
❤️344/657 💧445/445
🔋10/12 ⏰42 min 👣11
🪙577 💰9070
🎒Bag (3) /bag
Last battle info /report

Position: G 3#2
🌻Fields: 20🐑

More info /hero`,
			want: domain.Player{
				Level:           30,
				Rank:            1125,
				CurrentExp:      47154,
				NextLevelExp:    54934,
				BasicsUpdatedAt: time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "hero",
			msg: `🇲🇴M9co PakoB
🏅Level: 30 6.24%
📖Exp: 47154/54934
❤️345/657 💧445/445
🔋10/12 ⏰38 min 👣11

⚔️Rank: 1125
STR 8 DEX 13 VIT 12

🗡️Attack Force: 121
🌀Attack Speed: 339
⚡️Critical Rate: 90
💥Critical Force: 174
🦅Accuracy: 96
🥋Evasion: 36
🛡️Armor Score: 131
🥾Move Speed: 106

🪙577
🎒Bag (3) /bag
Expertise: 📘
Position: G 3#2
💡Skill points: 4 /skill
📚Schools /school
⚙️Settings /settings

🌟Noble 254 days /noble`,
			want: domain.Player{
				Level:           30,
				Rank:            1125,
				CurrentExp:      47154,
				NextLevelExp:    54934,
				BasicsUpdatedAt: time.Now(),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := domain.Player{}

			err := usecase.ParseBase(&player, tt.msg)
			if err != tt.wantErr {
				t.Errorf("err: expected %v, got %v", tt.wantErr, err)
			}

			assert(t, &tt.want, &player)
		})
	}
}
