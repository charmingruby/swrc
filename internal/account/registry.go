package account

import (
	"github.com/charmingruby/swrc/internal/account/domain/port"
	"github.com/charmingruby/swrc/internal/account/domain/repository"
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/account/infra/transport/grpc/server"
	"github.com/charmingruby/swrc/internal/common/infra/security"
	"google.golang.org/grpc"
)

func NewService(
	accountRepository repository.AccountRepository,
	hashAdapter port.HashPort,
) usecase.AccountUseCase {
	return usecase.NewAccountUseCaseRegistry(accountRepository, hashAdapter)
}

func NewGRPCHandler(srv *grpc.Server, accountSvc usecase.AccountUseCase, tokenSvc security.TokenService) {
	server.NewAccountGRPCServerHandler(accountSvc, tokenSvc, srv).Register()
}
