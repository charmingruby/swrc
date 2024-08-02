package server

import (
	"github.com/charmingruby/swrc/proto/pb"
	grpcLib "google.golang.org/grpc"
)

func NewAccountGRPCServerHandler(server *grpcLib.Server) *AccountGRPCServerHandler {
	return &AccountGRPCServerHandler{
		server: server,
	}
}

type AccountGRPCServerHandler struct {
	server *grpcLib.Server
}

func (h *AccountGRPCServerHandler) Register() {
	accountSvc := newAccountServiceGRPCServerHandler()

	pb.RegisterAccountServiceServer(h.server, accountSvc)
}
