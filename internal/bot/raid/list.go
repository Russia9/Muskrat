package raid

import (
	"context"
	"fmt"
	"gopkg.in/telebot.v3"
)

func (m *Module) List(c telebot.Context) error {
	raids, err := m.raid.List(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Raids:", raids)
	return nil
}
