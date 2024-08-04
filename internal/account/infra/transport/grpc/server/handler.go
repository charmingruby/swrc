package server

import (
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc"
)

func NewAccountGRPCServerHandler(
	accountSvc usecase.AccountUseCase,
	tokenSvc auth.TokenService,
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
	tokenService   auth.TokenService
}

func (h *AccountGRPCServerHandler) Register() {
	accountSvc := h.newAccountServiceGRPCServerHandler()
	pb.RegisterAccountServiceServer(h.server, accountSvc)
}
