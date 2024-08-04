package interceptor

import (
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"github.com/charmingruby/swrc/pkg/jwt"
)

func NewGRPCInterceptor(
	tokenSvc jwt.JWTService,
	authBypassMethods map[string]bool,
	rbacEnsuredMethods map[string][]string,
) GRPCInterceptor {
	return GRPCInterceptor{
		tokenService:       &tokenSvc,
		authBypassMethods:  authBypassMethods,
		rbacEnsuredMethods: rbacEnsuredMethods,
	}
}

type GRPCInterceptor struct {
	tokenService       auth.TokenService
	authBypassMethods  map[string]bool
	rbacEnsuredMethods map[string][]string
}
