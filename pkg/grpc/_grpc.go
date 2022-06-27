package grpc

import (
	"net"

	grpc_zerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/rs/zerolog/log"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
			),
			logging.UnaryServerInterceptor(grpc_zerolog.InterceptorLogger(log.Logger)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
			),
			logging.StreamServerInterceptor(grpc_zerolog.InterceptorLogger(log.Logger)),
		),
	)
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic().Err(err).Msg("failed to start tcp listener")
	}

	log.Info().Str("endpoint", addr).Msg("starting gRPC listener")

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Panic().Err(err).Msg("failed to start gRPC listener")
	}
}
