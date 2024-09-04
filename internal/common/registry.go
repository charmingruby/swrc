package common

import (
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/server"
	"google.golang.org/grpc"
)

func NewGRPCHandler(srv *grpc.Server) {
	server.NewCommonGRPCServerHandler(srv).Register()
}
