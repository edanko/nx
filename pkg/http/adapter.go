package http

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Adapter is http server app adapter
type Adapter struct {
	httpServer *http.Server
}

// NewAdapter provides new primary HTTP adapter
func NewAdapter(httpServer *http.Server) *Adapter {
	return &Adapter{
		httpServer: httpServer,
	}
}

// Start starts http application adapter
func (adapter *Adapter) Start(ctx context.Context) error {
	adapter.httpServer.BaseContext = func(_ net.Listener) context.Context { return ctx }

	log.Info().Str("endpoint", adapter.httpServer.Addr).Msg("starting HTTP listener")
	if err := adapter.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Panic().Err(err).Msg("failed to start HTTP server")
		return err
	}
	return nil
}

// Stop stops http application adapter
func (adapter *Adapter) Stop(ctx context.Context) error {
	return adapter.httpServer.Shutdown(ctx)
}
