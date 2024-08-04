package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.TrimPrefix(authHeader[0], "Bearer ")

	payload, err := i.tokenService.ValidateToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	ctx = context.WithValue(ctx, accountIDKey, payload.AccountID)
	ctx = context.WithValue(ctx, roleKey, payload.Role)
	ctx = context.WithValue(ctx, isValidKey, payload.IsValid)
	ctx = context.WithValue(ctx, verifiedKey, payload.Verified)

	return handler(ctx, req)
}
