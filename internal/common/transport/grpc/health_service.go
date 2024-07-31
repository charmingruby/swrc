package grpc

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/transport/grpc/mapper"
	"github.com/charmingruby/swrc/proto/pb"
)

func newGRPCHealthServiceHandler() *gRPCHealthServiceHandler {
	return &gRPCHealthServiceHandler{}
}

type gRPCHealthServiceHandler struct {
	pb.UnimplementedHealthServiceServer
}

func (h *gRPCHealthServiceHandler) HealthCheck(ctx context.Context, r *pb.PingMessage) (*pb.PingMessage, error) {
	msg := mapper.PingPongRequestGRPCToObj(r)
	msg.Greeting = msg.Greeting + " received"
	res := mapper.PingPongReplyObjToGRPC(msg)
	return res, nil
}
