package interceptor

import (
	"github.com/charmingruby/swrc/internal/common/infra/security"
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
	tokenService       security.TokenService
	authBypassMethods  map[string]bool
	rbacEnsuredMethods map[string][]string
}
