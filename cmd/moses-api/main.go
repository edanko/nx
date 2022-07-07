package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/edanko/nx/cmd/moses-api/internal/adapters/ent"
	"github.com/edanko/nx/cmd/moses-api/internal/adapters/ent/migrate"
	"github.com/edanko/nx/cmd/moses-api/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	loggerOutput := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339Nano,
	}
	log.Logger = zerolog.New(loggerOutput).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Str("app", "moses-api").
		Logger()

	cfg := config.GetConfig()

	if cfg.App.Environment == "development" {
		log.Logger = log.Logger.Level(zerolog.DebugLevel)
	}

	db, err := sql.Open("pgx", fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Database, cfg.DB.User, cfg.DB.Password, cfg.DB.SSLMode),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to database")
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(i ...any) {
			log.Debug().Str("message", fmt.Sprint(i...)).Msg("")
		}),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to database")
	}
	defer client.Close()

	if cfg.App.Environment == "development" {
		client = client.Debug()
	}

	err = client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

}
