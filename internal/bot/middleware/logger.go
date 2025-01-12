package middleware

import (
	"time"

	"github.com/Russia9/Muskrat/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
)

func (m *Middleware) Logger(next telebot.HandlerFunc) telebot.HandlerFunc {
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

			// Send error message
			tmpl := errorTmpl{
				Time:  time.Now().String(),
				User:  c.Sender().ID,
				Error: err.Error(),
			}
			c.Send(m.layout.Text(c, "error", tmpl))
			devChat, err := c.Bot().ChatByID(utils.DeveloperID)
			if err == nil {
				c.Bot().Send(devChat, m.layout.TextLocale("en", "error_dev", tmpl))
			}
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

type errorTmpl struct {
	Time  string
	User  int64
	Error string
}
