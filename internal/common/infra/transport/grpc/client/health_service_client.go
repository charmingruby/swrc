package client

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/contract"
)

func (c *CommonClientHandler) HealthCheck(ctx context.Context, in contract.PingMessage) (contract.PingMessage, error) {
	greq := contract.PingMessageObjToGRPC(in)

	gres, err := c.healthServiceClient.HealthCheck(ctx, &greq)
	if err != nil {
		return contract.PingMessage{}, err
	}

	res := contract.PingMessageGRPCToObj(gres)

	return res, nil
}
