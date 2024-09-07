package review_rpc

import (
	"github.com/charmingruby/swrc/internal/review/domain/usecase"
	"github.com/charmingruby/swrc/proto/pb"
)

func NewReviewGRPCService(
	reviewService usecase.ReviewUseCase,
) *ReviewGRPCService {
	return &ReviewGRPCService{
		reviewService: reviewService,
	}
}

type ReviewGRPCService struct {
	pb.UnimplementedReviewServiceServer
	reviewService usecase.ReviewUseCase
}
