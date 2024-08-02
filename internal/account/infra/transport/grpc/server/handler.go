package server

import (
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/account/infra/security"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc"
)

func NewAccountGRPCServerHandler(
	accountSvc usecase.AccountUseCase,
	tokenSvc security.TokenService,
	server *grpc.Server) *AccountGRPCServerHandler {
	return &AccountGRPCServerHandler{
		server:         server,
		accountService: accountSvc,
		tokenService:   tokenSvc,
	}
}

type AccountGRPCServerHandler struct {
	server         *grpc.Server
	accountService usecase.AccountUseCase
	tokenService   security.TokenService
}

func (h *AccountGRPCServerHandler) Register() {
	accountSvc := h.newAccountServiceGRPCServerHandler()
	pb.RegisterAccountServiceServer(h.server, accountSvc)
}
