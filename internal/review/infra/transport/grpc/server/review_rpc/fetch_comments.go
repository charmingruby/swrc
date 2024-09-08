package review_rpc

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *ReviewGRPCService) FetchComments(ctx context.Context, req *pb.FetchCommentsRequest) (*pb.FetchCommentsReply, error) {
	input := dto.FetchCommentsInputDTO{
		ID:              req.Id,
		SnippetTopicID:  req.SnippetTopicId,
		AccountID:       req.AccountId,
		ParentCommentID: req.ParentCommentId,
	}

	output, err := h.reviewService.FetchCommentsUseCase(input)
	if err != nil {
		return nil, grpc.NewInternalErr(err)
	}

	var parsedComments []*pb.Comment

	if len(output.Comments) != 0 {
		for _, c := range output.Comments {
			parsedComments = append(parsedComments, &pb.Comment{
				Id:              c.ID,
				Content:         c.Content,
				ParentCommentId: c.ParentCommentID,
				SnippetTopicId:  c.SnippetTopicID,
				AccountId:       c.AccountID,
				CreatedAt:       timestamppb.New(c.CreatedAt),
			})
		}
	}

	res := pb.FetchCommentsReply{
		Comments: parsedComments,
	}

	return &res, nil
}
