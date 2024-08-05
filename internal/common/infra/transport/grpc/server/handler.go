package server

import (
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/server/health_service"
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
	healthSvc := health_service.NewHealthGRPCService()
	pb.RegisterHealthServiceServer(h.server, healthSvc)
}
