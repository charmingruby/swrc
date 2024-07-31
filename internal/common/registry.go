package common

import (
	commonGRPC "github.com/charmingruby/swrc/internal/common/transport/grpc"
	"google.golang.org/grpc"
)

func NewCommonGRPCHandlerSetup(server *grpc.Server) {
	commonGRPC.NewCommonGRPCHandler(server).Register()
}
