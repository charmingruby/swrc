package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	accountIDKey contextKey = "accountID"
	roleKey      contextKey = "role"
	isValidKey   contextKey = "isValid"
	verifiedKey  contextKey = "verified"
)

func (i *GRPCInterceptor) AuthInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	if i.authBypassMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	payload, ctx, err := i.retrieveTokenFromMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	ctx = i.savePayloadAtCtx(ctx, payload)

	return handler(ctx, req)
}
