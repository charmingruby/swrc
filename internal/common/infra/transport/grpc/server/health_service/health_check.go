package health_service

import (
	"context"

	"github.com/charmingruby/swrc/proto/pb"
)

func (h *HealthGRPCService) HealthCheck(ctx context.Context, r *pb.PingMessage) (*pb.PingMessage, error) {
	res := pb.PingMessage{
		Greeting: r.Greeting,
	}
	return &res, nil
}
