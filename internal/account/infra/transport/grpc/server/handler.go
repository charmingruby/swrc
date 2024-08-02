package server

import (
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc"
)

func NewAccountGRPCServerHandler(
	accountSvc usecase.AccountUseCase,
	server *grpc.Server) *AccountGRPCServerHandler {
	return &AccountGRPCServerHandler{
		server:         server,
		accountService: accountSvc,
	}
}

type AccountGRPCServerHandler struct {
	server         *grpc.Server
	accountService usecase.AccountUseCase
}

func (h *AccountGRPCServerHandler) Register() {
	accountSvc := newAccountServiceGRPCServerHandler(h.accountService)
	pb.RegisterAccountServiceServer(h.server, accountSvc)
}
