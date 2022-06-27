package http

import (
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"

	httpmiddleware "github.com/edanko/nx/pkg/http/middleware"
	"github.com/edanko/nx/pkg/logger"
)

// func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
// 	apiRouter := chi.NewRouter()
// 	SetMiddlewares(apiRouter)

// 	mainRouter := chi.NewRouter()
// 	mainRouter.Get("/health", healthHandler())
// 	mainRouter.Get("/readiness", readinessHandler())
// 	mainRouter.Mount("/v1", createHandler(apiRouter))

// 	log.Info().Str("endpoint", addr).Msg("starting HTTP listener")

// 	err := http.ListenAndServe(addr, mainRouter)
// 	if err != nil {
// 		log.Panic().Err(err).Msg("failed to start HTTP server")
// 	}
// }

func SetMiddlewares(router *chi.Mux) {
	router.Use(httpmiddleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logger.NewStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)
	// addAuthMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

// func addAuthMiddleware(router *chi.Mux) {
// 	router.Use(auth.HttpMockMiddleware)

// if mockAuth, _ := strconv.ParseBool(os.Getenv("MOCK_AUTH")); mockAuth {
// 	router.Use(auth.HttpMockMiddleware)
// 	return
// }

// var opts []option.ClientOption
// if file := os.Getenv("SERVICE_ACCOUNT_FILE"); file != "" {
// 	opts = append(opts, option.WithCredentialsFile(file))
// }

// config := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT")}
// firebaseApp, err := firebase.NewApp(context.Background(), config, opts...)
// if err != nil {
// 	logrus.Fatalf("error initializing app: %v\n", err)
// }

// authClient, err := firebaseApp.Auth(context.Background())
// if err != nil {
// 	logrus.WithError(err).Fatal("Unable to create firebase Auth client")
// }

// router.Use(auth.FirebaseHttpMiddleware{authClient}.Middleware)
// }

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}
