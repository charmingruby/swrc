package review_rpc

import (
	"context"

	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *ReviewGRPCService) FetchSnippets(ctx context.Context, req *pb.FetchSnippetsRequest) (*pb.FetchSnippetsReply, error) {
	input := dto.FetchSnippetInputDTO{
		SnippetTopicID: req.SnippetTopicId,
	}

	output, err := h.reviewService.FetchSnippetsUseCase(input)
	if err != nil {
		return nil, grpc.NewInternalErr(err)
	}

	var parsedSnippets []*pb.Snippet

	if len(output.Snippets) != 0 {
		for _, s := range output.Snippets {
			parsedSnippets = append(parsedSnippets, &pb.Snippet{
				Id:          s.ID,
				Version:     int64(s.Version),
				CodeSnippet: s.CodeSnippet,
				Message:     s.Message,
				TopicId:     s.TopicID,
				CreatedAt:   timestamppb.New(s.CreatedAt),
			})
		}
	}

	res := pb.FetchSnippetsReply{
		Snippets: parsedSnippets,
	}

	return &res, nil
}
