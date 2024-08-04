package account

import (
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/account/infra/transport/grpc/server"
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"google.golang.org/grpc"
)

func NewAccountGRPCHandlerSetup(srv *grpc.Server, accountSvc usecase.AccountUseCase, tokenSvc auth.TokenService) {
	server.NewAccountGRPCServerHandler(accountSvc, tokenSvc, srv).Register()
}
