package account

import (
	"github.com/charmingruby/swrc/internal/account/transport/grpc/server"
	"google.golang.org/grpc"
)

func NewAccountGRPCHandlerSetup(srv *grpc.Server) {
	server.NewAccountGRPCServerHandler(srv).Register()
}
