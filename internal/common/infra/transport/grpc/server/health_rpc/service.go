package health_rpc

import "github.com/charmingruby/swrc/proto/pb"

func NewHealthGRPCService() *HealthGRPCService {
	return &HealthGRPCService{}
}

type HealthGRPCService struct {
	pb.UnimplementedHealthServiceServer
}
