package decorator

import (
	"context"
	"errors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryAuthInterceptor returns a grpc.UnaryServerInterceptor that enforces a token.
func UnaryAuthInterceptor(validToken string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata")
		}
		tokens := md["authorization"]
		if len(tokens) == 0 || tokens[0] != validToken {
			return nil, errors.New("unauthorized")
		}
		// Delegate to the actual handler
		return handler(ctx, req)
	}
}

func RunAuthServer() {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryAuthInterceptor("s3cr3t-token")),
	)
	// register your services...
	lis, _ := net.Listen("tcp", ":50051")
	server.Serve(lis)
}
