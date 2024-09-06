package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *GRPCInterceptor) RBACInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	roles, ok := i.rbacEnsuredMethods[info.FullMethod]
	if !ok {
		return handler(ctx, req)
	}

	payload, ctx, err := i.retrieveTokenFromMetadata(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if isRoleIncluded := i.isRoleIncluded(payload.Role, roles); !isRoleIncluded {
		return nil, status.Errorf(codes.Unauthenticated, "don't have needed permissions")
	}

	ctx = i.savePayloadAtCtx(ctx, payload)

	return handler(ctx, req)
}

func (i *GRPCInterceptor) isRoleIncluded(accountRole string, roles []string) bool {
	for _, r := range roles {
		if r == accountRole {
			return true
		}
	}

	return false
}
