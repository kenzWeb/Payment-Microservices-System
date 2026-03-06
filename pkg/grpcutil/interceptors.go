package grpcutil

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/user/payment-microservices/pkg/logger"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if tid := md.Get("trace_id"); len(tid) > 0 {
				ctx = context.WithValue(ctx, logger.TraceIDKey, tid[0])
			}
		}
		return handler(ctx, req)
	}
}

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if tid, ok := ctx.Value(logger.TraceIDKey).(string); ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "trace_id", tid)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
