package server

import (
	"github.com/charmingruby/swrc/proto/pb"
	grpcLib "google.golang.org/grpc"
)

func NewAccountGRPCServerHandler(server *grpcLib.Server) *AccountGRPCHandler {
	return &AccountGRPCHandler{
		server: server,
	}
}

type AccountGRPCHandler struct {
	server *grpcLib.Server
}

func (h *AccountGRPCHandler) Register() {
	accountSvc := NewGRPCAccountServiceHandler()

	pb.RegisterAccountServiceServer(h.server, accountSvc)
}
