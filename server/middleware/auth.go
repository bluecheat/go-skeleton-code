package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func AuthUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	// retrieve metadata from context
	//md, ok := metadata.FromIncomingContext(ctx)
	//logger.Info(md, ok, info.FullMethod, ctx, req)
	// validate 'authorization' metadata
	// like headers, the value is an slice []string
	//uid, err := MyValidationFunc(md["authorization"])
	//if err != nil {
	//	return nil, status.Errorf(codes.Unauthenticated, "authentication required")
	//}

	// add user ID to the context
	//newCtx := context.WithValue(ctx, "user_id", uid)

	// handle scopes?
	return handler(ctx, req)
}
