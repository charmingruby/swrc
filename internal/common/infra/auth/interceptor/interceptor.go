package interceptor

import (
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"github.com/charmingruby/swrc/pkg/jwt"
)

func NewGRPCInterceptor(tokenSvc jwt.JWTService, authBypassMethods map[string]bool) GRPCInterceptor {
	return GRPCInterceptor{
		tokenService:      &tokenSvc,
		authBypassMethods: authBypassMethods,
	}
}

type GRPCInterceptor struct {
	tokenService      auth.TokenService
	authBypassMethods map[string]bool
}
