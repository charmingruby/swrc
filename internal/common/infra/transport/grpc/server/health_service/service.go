package health_service

import "github.com/charmingruby/swrc/proto/pb"

func NewHealthGRPCService() *HealthGRPCService {
	return &HealthGRPCService{}
}

type HealthGRPCService struct {
	pb.UnimplementedHealthServiceServer
}
