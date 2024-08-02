package account

import (
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/account/infra/security"
	"github.com/charmingruby/swrc/internal/account/infra/transport/grpc/server"
	"google.golang.org/grpc"
)

func NewAccountGRPCHandlerSetup(srv *grpc.Server, accountSvc usecase.AccountUseCase, tokenSvc security.TokenService) {
	server.NewAccountGRPCServerHandler(accountSvc, tokenSvc, srv).Register()
}
