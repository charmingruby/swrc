package server

import (
	"github.com/charmingruby/swrc/internal/review/domain/usecase"
	"github.com/charmingruby/swrc/internal/review/infra/transport/grpc/server/review_rpc"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc"
)

func NewReviewGRPCServerHandler(
	reviewSvc usecase.ReviewUseCase,
	server *grpc.Server,
) *ReviewGRPCServerHandler {
	return &ReviewGRPCServerHandler{
		server:        server,
		reviewService: reviewSvc,
	}
}

type ReviewGRPCServerHandler struct {
	server        *grpc.Server
	reviewService usecase.ReviewUseCase
}

func (h *ReviewGRPCServerHandler) Register() {
	reviewSvc := review_rpc.NewReviewGRPCService(h.reviewService)
	pb.RegisterReviewServiceServer(h.server, reviewSvc)
}
