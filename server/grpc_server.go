package server

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
	"skeleton-code/components"
	"skeleton-code/config"
	"skeleton-code/logger"
	"skeleton-code/proto/generated"
	"skeleton-code/server/handlers"
	"skeleton-code/server/middleware"
)

func GRPCServer(lifecycle fx.Lifecycle, conf *config.Config, ctx components.Context) {
	var ip string
	if conf.Host == "" {
		ip = "localhost"
	} else {
		ip = conf.Host
	}

	host := fmt.Sprintf("%s:%s", ip, conf.RPCPort)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		logger.Error("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(), //서버 리커버리
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),
			middleware.LoggerUnaryInterceptor,
			middleware.AuthUnaryInterceptor,
			grpc_prometheus.UnaryServerInterceptor,
		),
	)

	vehicleHandler := handlers.NewVehicleHandler(ctx)
	memberHandler := handlers.NewMemberHandler(ctx)
	generated.RegisterVehicleServiceServer(s, vehicleHandler)
	generated.RegisterMemberServiceServer(s, memberHandler)

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("run grpc service", host)
				go s.Serve(lis)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				s.Stop()
				return nil
			},
		},
	)
}
