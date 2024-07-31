package grpc

import (
	"github.com/charmingruby/swrc/internal/common/transport/grpc/server"
	"github.com/charmingruby/swrc/proto/pb"
	grpcLib "google.golang.org/grpc"
)

func NewCommonGRPCHandler(server *grpcLib.Server) *CommonGRPCHandler {
	return &CommonGRPCHandler{
		server: server,
	}
}

type CommonGRPCHandler struct {
	server *grpcLib.Server
}

func (h *CommonGRPCHandler) Register() {
	healthSvc := server.NewGRPCHealthServiceHandler()

	pb.RegisterHealthServiceServer(h.server, healthSvc)
}
