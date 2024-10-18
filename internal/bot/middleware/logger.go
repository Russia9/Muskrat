package middleware

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
)

func Logger(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		// Start timer
		start := time.Now()

		// Process request
		err := next(c)

		// Create log entry
		var record *zerolog.Event
		if err == nil {
			record = log.Info()
		} else {
			record = log.Error().Err(err)
		}

		// Add log fields
		record.
			Int64("user", c.Sender().ID).
			Str("latency", time.Since(start).String())
		if c.Message().Text != "" {
			record.Str("text", c.Message().Text)
		}
		if c.Callback() != nil {
			record.Str("callback", c.Callback().Unique)
		}

		// Send log entry
		record.Send()

		return nil
	}
}
