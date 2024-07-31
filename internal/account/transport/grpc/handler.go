package grpc

import (
	"github.com/charmingruby/swrc/internal/account/transport/grpc/server"
	"github.com/charmingruby/swrc/proto/pb"
	grpcLib "google.golang.org/grpc"
)

func NewAccountGRPCHandler(server *grpcLib.Server) *AccountGRPCHandler {
	return &AccountGRPCHandler{
		server: server,
	}
}

type AccountGRPCHandler struct {
	server *grpcLib.Server
}

func (h *AccountGRPCHandler) Register() {
	accountSvc := server.NewGRPCAccountServiceHandler()

	pb.RegisterAccountServiceServer(h.server, accountSvc)
}
