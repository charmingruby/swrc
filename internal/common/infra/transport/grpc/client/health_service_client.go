package client

import (
	"context"

	"github.com/charmingruby/swrc/proto/pb"
)

func (c *CommonClientHandler) HealthCheck(ctx context.Context, req *pb.PingMessage) (*pb.PingMessage, error) {
	rep, err := c.healthServiceClient.HealthCheck(ctx, req)
	if err != nil {
		return nil, err
	}

	return rep, nil
}
