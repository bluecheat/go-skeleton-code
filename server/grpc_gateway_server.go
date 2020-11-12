package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net/http"
	"skeleton-code/config"
	"skeleton-code/logger"
	"skeleton-code/proto/generated"
)

func APIServer(lifecycle fx.Lifecycle, conf *config.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	err := generated.RegisterVehicleServiceHandlerFromEndpoint(ctx, mux, "localhost:"+conf.RPCPort, opts)
	if err != nil {
		logger.Error(err)
	}
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				var ip string
				if conf.Host == "" {
					ip = "localhost"
				} else {
					ip = conf.Host
				}
				host := fmt.Sprintf("%s:%s", ip, conf.Port)
				logger.Info("run http restful service", host)
				go http.ListenAndServe(host, mux)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				cancel()
				return nil
			},
		},
	)
}
