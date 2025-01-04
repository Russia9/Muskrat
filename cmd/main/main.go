package main

import (
	"context"
	"os"
	"strconv"

	playerRepo "github.com/Russia9/Muskrat/internal/player/repository/mongo"
	playerUsecase "github.com/Russia9/Muskrat/internal/player/usecase"
	squadRepo "github.com/Russia9/Muskrat/internal/squad/repository/mongo"
	squadUsecae "github.com/Russia9/Muskrat/internal/squad/usecase"

	"github.com/Russia9/Muskrat/internal/bot"
	"github.com/Russia9/Muskrat/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// DB Connection
	log.Debug().Msg("DB Connection")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal().Err(err).Msg("DB Connection")
	}
	db := client.Database(utils.GetEnv("MONGO_DB", "muskrat"))

	// DB Ping
	log.Debug().Msg("DB Ping")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("DB Ping")
	}

	// Repository creation
	log.Debug().Msg("Repository creation")
	playerRepo := playerRepo.NewPlayerRepo(db)
	squadRepo := squadRepo.NewSquadRepo(db)

	// Usecase creation
	log.Debug().Msg("Usecase creation")
	playerUC := playerUsecase.NewPlayerUsecase(playerRepo)
	squadUC := squadUsecae.NewSquadUsecase(squadRepo, playerRepo)

	// Bot
	log.Trace().Msg("Layout loading")
	l, err := layout.New("assets/layout/layout.yml")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load layout")
	}
	log.Trace().Msg("Bot creation")
	tb, err := telebot.NewBot(l.Settings())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bot")
	}
	b := bot.NewBot(tb, l, playerUC, squadUC)

	// Start bot
	log.Info().Msg("Starting bot")
	b.StartBlocking()
}
