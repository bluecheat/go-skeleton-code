package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"skeleton-code/logger"
)

func AuthUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	// retrieve metadata from context
	m, ok := metadata.FromIncomingContext(ctx)
	logger.Info(m, info.FullMethod)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "authentication required")
	}

	// add user ID to the context
	//newCtx := context.WithValue(ctx, "user_id", uid)

	// handle scopes?
	return handler(ctx, req)
}
