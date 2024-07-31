package account

import (
	accountGRPC "github.com/charmingruby/swrc/internal/account/transport/grpc"
	"google.golang.org/grpc"
)

func NewAccountGRPCHandlerSetup(server *grpc.Server) {
	accountGRPC.NewAccountGRPCHandler(server).Register()
}
