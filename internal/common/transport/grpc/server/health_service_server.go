package server

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/transport/grpc/mapper"
	"github.com/charmingruby/swrc/proto/pb"
)

func NewGRPCHealthServiceHandler() *GRPCHealthServiceHandler {
	return &GRPCHealthServiceHandler{}
}

type GRPCHealthServiceHandler struct {
	pb.UnimplementedHealthServiceServer
}

func (h *GRPCHealthServiceHandler) HealthCheck(ctx context.Context, r *pb.PingMessage) (*pb.PingMessage, error) {
	msg := mapper.PingPongRequestGRPCToObj(r)
	msg.Greeting = msg.Greeting + " received"
	res := mapper.PingPongReplyObjToGRPC(msg)
	return res, nil
}
