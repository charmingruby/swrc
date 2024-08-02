package server

import (
	"github.com/charmingruby/swrc/proto/pb"
	grpcLib "google.golang.org/grpc"
)

func NewCommonGRPCServerHandler(server *grpcLib.Server) *CommonGRPCHandler {
	return &CommonGRPCHandler{
		server: server,
	}
}

type CommonGRPCHandler struct {
	server *grpcLib.Server
}

func (h *CommonGRPCHandler) Register() {
	healthSvc := newHealthServiceServerHandler()
	pb.RegisterHealthServiceServer(h.server, healthSvc)
}
