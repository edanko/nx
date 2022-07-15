package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ThreeDotsLabs/watermill-jetstream/pkg/jetstream"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	watermillmiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/microcosm-cc/bluemonday"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/edanko/nx/cmd/launch-api/internal/adapters"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/migrate"
	"github.com/edanko/nx/cmd/launch-api/internal/app"
	"github.com/edanko/nx/cmd/launch-api/internal/app/commands"
	"github.com/edanko/nx/cmd/launch-api/internal/app/events"
	"github.com/edanko/nx/cmd/launch-api/internal/app/queries"
	"github.com/edanko/nx/cmd/launch-api/internal/config"
	"github.com/edanko/nx/cmd/launch-api/internal/ports"
	"github.com/edanko/nx/pkg/application"
	httputils "github.com/edanko/nx/pkg/http"
	"github.com/edanko/nx/pkg/http/handlers"
	"github.com/edanko/nx/pkg/logger"
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
		Str("app", "launch-api").
		Logger()

	cfg := config.GetConfig()

	if cfg.App.Environment == "development" {
		log.Logger = log.Logger.Level(zerolog.DebugLevel)
	}

	watermillLogger := logger.NewZerologLoggerAdapter(log.Logger)

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
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	// watermill
	natsURL := fmt.Sprintf("%s:%d", cfg.NATS.Host, cfg.NATS.Port)

	exactlyOnceDelivery := true

	natsOptions := []nats.Option{
		nats.RetryOnFailedConnect(true),
		nats.Timeout(30 * time.Second),
		nats.ReconnectWait(1 * time.Second),
	}

	c, err := nats.Connect(natsURL, natsOptions...)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connecting to nats")
	}
	defer c.Close()

	_, err = c.JetStream()
	if err != nil {
		log.Fatal().Err(err).Msg("jetstream error")
	}

	jetstreamMarshalerUnmarshaler := jetstream.GobMarshaler{}

	jetstreamOptions := make([]nats.JSOpt, 0)

	jetstreamPublisher, err := jetstream.NewPublisher(
		jetstream.PublisherConfig{
			URL:              natsURL,
			Marshaler:        jetstreamMarshalerUnmarshaler,
			NatsOptions:      natsOptions,
			JetstreamOptions: jetstreamOptions,
			AutoProvision:    true,
			TrackMsgId:       exactlyOnceDelivery,
		},
		watermillLogger,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to creating jetstream publisher")
	}
	defer jetstreamPublisher.Close()

	subscribeOptions := []nats.SubOpt{
		nats.DeliverAll(),
		nats.AckExplicit(),
	}
	jetstreamSubscriber, err := jetstream.NewSubscriber(
		jetstream.SubscriberConfig{
			URL:              natsURL,
			Unmarshaler:      jetstreamMarshalerUnmarshaler,
			QueueGroup:       "launch",
			DurableName:      "durable-name",
			SubscribersCount: 10,
			AckWaitTimeout:   30 * time.Second,
			NatsOptions:      natsOptions,
			SubscribeOptions: subscribeOptions,
			JetstreamOptions: jetstreamOptions,
			CloseTimeout:     30 * time.Second,
			AutoProvision:    true,
			AckSync:          exactlyOnceDelivery,
		},
		watermillLogger,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to creating jetstream commands subscriber")
	}
	defer jetstreamSubscriber.Close()

	router, err := message.NewRouter(
		message.RouterConfig{
			CloseTimeout: 30 * time.Second,
		},
		watermillLogger,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to creating router")
	}
	defer router.Close()

	router.AddMiddleware(watermillmiddleware.Recoverer)

	commandEventMarshaler := cqrs.JSONMarshaler{}
	generateCommandsTopic := func(commandName string) string {
		return strings.Replace(commandName, ".", "_", 1)
	}
	commandsSubscriberConstructor := func(handlerName string) (message.Subscriber, error) {
		return jetstreamSubscriber, nil
	}
	generateEventsTopic := func(eventName string) string {
		return strings.Replace(eventName, ".", "_", 1)
	}
	eventsSubscriberConstructor := func(_ string) (message.Subscriber, error) {
		return jetstream.NewSubscriber(
			jetstream.SubscriberConfig{
				URL:              natsURL,
				Unmarshaler:      jetstreamMarshalerUnmarshaler,
				DurableName:      "durable-name",
				AckWaitTimeout:   30 * time.Second,
				NatsOptions:      natsOptions,
				SubscribeOptions: subscribeOptions,
				JetstreamOptions: jetstreamOptions,
				CloseTimeout:     30 * time.Second,
				AutoProvision:    true,
				AckSync:          exactlyOnceDelivery,
			},
			watermillLogger,
		)
	}

	commandBus, err := cqrs.NewCommandBus(
		jetstreamPublisher,
		generateCommandsTopic,
		commandEventMarshaler,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to creating commands bus")
	}

	eventBus, err := cqrs.NewEventBus(
		jetstreamPublisher,
		generateEventsTopic,
		commandEventMarshaler,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to creating events bus")
	}

	sanitizer := bluemonday.UGCPolicy()

	kindsRepository := adapters.NewKindRepository(client.Kind)

	commandHandlers := []cqrs.CommandHandler{
		commands.NewCreateKindHandler(eventBus, kindsRepository, sanitizer),
		commands.NewDeleteKindHandler(eventBus, kindsRepository),
		commands.NewChangeKindNameHandler(eventBus, kindsRepository, sanitizer),
		commands.NewChangeKindDescriptionHandler(eventBus, kindsRepository, sanitizer),
		commands.NewMakeKindPublishedHandler(eventBus, kindsRepository),
		commands.NewMakeKindDraftHandler(eventBus, kindsRepository),
	}

	commandProcessor, err := cqrs.NewCommandProcessor(
		commandHandlers,
		generateCommandsTopic,
		commandsSubscriberConstructor,
		commandEventMarshaler,
		watermillLogger,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to creating commands processor")
	}

	if err := commandProcessor.AddHandlersToRouter(router); err != nil {
		log.Fatal().Err(err).Msg("failed to adding commands handlers to router")
	}

	eventHandlers := []cqrs.EventHandler{
		events.KindCreatedHandler{},
		events.KindDeletedHandler{},
		events.KindNameChangedHandler{},
		events.KindDescriptionChangedHandler{},
		events.KindMadePublishedHandler{},
		events.KindMadeDraftHandler{},
	}

	eventProcessor, err := cqrs.NewEventProcessor(
		eventHandlers,
		generateEventsTopic,
		eventsSubscriberConstructor,
		commandEventMarshaler,
		watermillLogger,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to creating events processor")
	}

	if err := eventProcessor.AddHandlersToRouter(router); err != nil {
		log.Fatal().Err(err).Msg("failed to adding events handlers to router")
	}

	app := app.Application{
		CommandBus: commandBus,
		Queries: app.Queries{
			ListKinds:     queries.NewListKindsHandler(kindsRepository),
			GetKind:       queries.NewGetKindHandler(kindsRepository),
			GetKindByName: queries.NewGetKindByNameHandler(kindsRepository),
		},
	}

	go func(ctx context.Context) {
		err := router.Run(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to run router")
		}
	}(ctx)
	<-router.Running()
	// watermill

	// router
	apiRouter := chi.NewRouter()
	httputils.SetMiddlewares(apiRouter)

	mainRouter := chi.NewRouter()
	mainRouter.Get("/version", handlers.BuildVersionHandler())
	mainRouter.Get("/health", handlers.BuildHealthHandler())
	// mainRouter.Get("/readiness", handlers.BuildReadinessHandler())
	mainRouter.Mount("/v1", ports.HandlerFromMux(
		ports.NewHTTPServer(app),
		apiRouter,
	))
	// router

	a := application.New()
	a.AddAdapters(
		httputils.NewAdapter(
			&http.Server{
				Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
				ReadTimeout:  cfg.HTTP.ReadTimeout,
				WriteTimeout: cfg.HTTP.WriteTimeout,
				IdleTimeout:  cfg.HTTP.IdleTimeout,
				Handler:      mainRouter,
			},
		),
	)

	if cfg.App.Environment == "development" {
		// a.AddAdapters(
		// 	application.NewDebugAdapter(
		// 		fmt.Sprintf("%s:%d", cfg.Debug.Host, cfg.Debug.Port),
		// 	),
		// )
	}

	a.WithShutdownTimeout(cfg.App.ShutdownTimeout)
	a.Run(ctx)
}
