package account

import (
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/account/transport/grpc/server"
	"google.golang.org/grpc"
)

func NewAccountGRPCHandlerSetup(srv *grpc.Server, accountSvc usecase.AccountUseCase) {
	server.NewAccountGRPCServerHandler(accountSvc, srv).Register()
}
