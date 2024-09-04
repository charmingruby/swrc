package review_service

import (
	"context"

	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *ReviewGRPCService) CreateSnippetTopic(ctx context.Context, in *pb.CreateSnippetTopicRequest) (*emptypb.Empty, error) {
	return nil, nil
}
