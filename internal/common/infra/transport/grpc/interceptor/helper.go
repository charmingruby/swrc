package interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmingruby/swrc/internal/common/infra/security"
	"google.golang.org/grpc/metadata"
)

type contextKey string

const (
	AccountIDKey contextKey = "accountID"
	RoleKey      contextKey = "role"
	IsValidKey   contextKey = "isValid"
	VerifiedKey  contextKey = "verified"
)

func (i *GRPCInterceptor) retrieveTokenFromMetadata(
	ctx context.Context,
) (security.TokenPayload, context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return security.TokenPayload{}, ctx, fmt.Errorf("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return security.TokenPayload{}, ctx, fmt.Errorf("authorization token is not provided")
	}

	token := strings.TrimPrefix(authHeader[0], "Bearer ")

	payload, err := i.tokenService.ValidateToken(token)
	if err != nil {
		return security.TokenPayload{}, ctx, fmt.Errorf("invalid token: %v", err)
	}

	return payload, ctx, nil
}

func (i *GRPCInterceptor) savePayloadAtCtx(ctx context.Context, payload security.TokenPayload) context.Context {
	ctx = context.WithValue(ctx, AccountIDKey, payload.AccountID)
	ctx = context.WithValue(ctx, RoleKey, payload.Role)
	ctx = context.WithValue(ctx, IsValidKey, payload.IsValid)
	ctx = context.WithValue(ctx, VerifiedKey, payload.Verified)

	return ctx
}
