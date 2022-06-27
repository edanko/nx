package grpc

import (
	"context"
	"net"

	"github.com/rs/zerolog/log"
	grpcmain "google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	grpchealthproto "google.golang.org/grpc/health/grpc_health_v1"
)

// Adapter is grpc server app adapter
type Adapter struct {
	name         string
	address      string
	server       *grpcmain.Server
	healthServer *grpchealth.Server
}

// NewAdapter provides new primary adapter
func NewAdapter(name, address string, server *grpcmain.Server) *Adapter {
	return &Adapter{
		name:         name,
		address:      address,
		server:       server,
		healthServer: grpchealth.NewServer(),
	}
}

// Start starts grpc application adapter
func (adapter *Adapter) Start(_ context.Context) error {
	grpchealthproto.RegisterHealthServer(adapter.server, adapter.healthServer)

	listener, err := net.Listen("tcp", adapter.address)
	if err != nil {
		log.Warn().Err(err).Msg("failed to start tcp listener")
		return err
	}

	log.Info().Str("endpoint", adapter.address).Msg("starting gRPC listener")

	adapter.healthServer.SetServingStatus(adapter.name, grpchealthproto.HealthCheckResponse_SERVING)

	return adapter.server.Serve(listener)
}

// Stop stops grpc application adapter
func (adapter *Adapter) Stop(_ context.Context) error {
	adapter.healthServer.SetServingStatus(
		adapter.name,
		grpchealthproto.HealthCheckResponse_NOT_SERVING,
	)

	adapter.server.GracefulStop()

	return nil
}
