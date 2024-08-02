package server

import (
	"context"

	"github.com/charmingruby/swrc/proto/pb"
)

func newHealthServiceServerHandler() *HealthServiceServerHandler {
	return &HealthServiceServerHandler{}
}

type HealthServiceServerHandler struct {
	pb.UnimplementedHealthServiceServer
}

func (h *HealthServiceServerHandler) HealthCheck(ctx context.Context, r *pb.PingMessage) (*pb.PingMessage, error) {
	res := pb.PingMessage{
		Greeting: r.Greeting,
	}
	return &res, nil
}
