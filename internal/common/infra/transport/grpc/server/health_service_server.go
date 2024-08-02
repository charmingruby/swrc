package server

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/contract"
	"github.com/charmingruby/swrc/proto/pb"
)

func newHealthServiceServerHandler() *HealthServiceServerHandler {
	return &HealthServiceServerHandler{}
}

type HealthServiceServerHandler struct {
	pb.UnimplementedHealthServiceServer
}

func (h *HealthServiceServerHandler) HealthCheck(ctx context.Context, r *pb.PingMessage) (*pb.PingMessage, error) {
	msg := contract.PingMessageGRPCToObj(r)
	msg.Greeting = msg.Greeting + " received"
	res := contract.PingMessageObjToGRPC(msg)
	return &res, nil
}
