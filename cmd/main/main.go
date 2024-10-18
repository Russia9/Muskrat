package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

func main() {
	// Log settings
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	pretty, err := strconv.ParseBool(os.Getenv("LOG_PRETTY"))
	if err != nil {
		pretty = false
	}
	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	switch os.Getenv("LOG_LEVEL") {
	case "DISABLED":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Bot
	log.Debug().Msg("Creating bot")
	log.Trace().Msg("Loading layout")
	parsedLayout, err := layout.New("assets/layout/layout.yml")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load layout")
	}
	log.Trace().Msg("Creating bot")
	bot, err := telebot.NewBot(parsedLayout.Settings())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bot")
	}

	fmt.Println(bot.Me.ID)
}
