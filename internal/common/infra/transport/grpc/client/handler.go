package client

import (
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc"
)

func NewCommonClientHandler(conn *grpc.ClientConn) *CommonClientHandler {
	return &CommonClientHandler{
		healthServiceClient: pb.NewHealthServiceClient(conn),
	}
}

type CommonClientHandler struct {
	healthServiceClient pb.HealthServiceClient
}
