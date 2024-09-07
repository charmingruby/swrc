package review_service

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *ReviewGRPCService) FetchSnippetTopics(ctx context.Context, req *pb.FetchSnippetTopicsRequest) (*pb.FetchSnippetTopicsReply, error) {
	input := dto.FetchSnippetTopicsInputDTO{
		ID:        req.Id,
		Status:    req.Status,
		AccountID: req.AccountId,
	}

	output, err := h.reviewService.FetchSnippetTopicsUseCase(input)
	if err != nil {
		return nil, grpc.NewInternalErr(err)
	}

	var parsedTopics []*pb.SnippetTopic

	if len(output.Topics) != 0 {
		for _, t := range output.Topics {
			parsedTopics = append(parsedTopics, &pb.SnippetTopic{
				Id:                t.ID,
				Title:             t.Title,
				Description:       t.Description,
				Status:            t.Status,
				CurrentVersion:    int64(t.CurrentVersion),
				BestAnswerId:      t.BestAnswerID,
				SnippetSolutionId: t.SnippetSolutionID,
				AccountId:         t.AccountID,
				CreatedAt:         timestamppb.New(t.CreatedAt),
			})
		}
	}

	res := pb.FetchSnippetTopicsReply{
		Topics: parsedTopics,
	}

	return &res, nil
}
