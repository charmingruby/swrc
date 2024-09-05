package review_service

import (
	"context"
	"errors"

	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/infra/auth/interceptor"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/internal/review/domain/dto"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *ReviewGRPCService) CreateSnippetTopic(ctx context.Context, req *pb.CreateSnippetTopicRequest) (*emptypb.Empty, error) {
	accountID := ctx.Value(interceptor.AccountIDKey).(string)
	if accountID == "" {
		return nil, grpc.NewInternalErr(errors.New("account id not found"))
	}

	input := dto.CreateSnippetTopicInputDTO{
		Title:       req.Title,
		Description: req.Description,
		AccountID:   accountID,
	}

	if err := h.reviewService.CreateSnippetTopicUseCase(input); err != nil {
		if notFoundErr, ok := err.(*core.ErrNotFound); ok {
			return nil, grpc.NewNotFoundErr(notFoundErr)
		}

		if validationErr, ok := err.(*core.ErrValidation); ok {
			return nil, grpc.NewValidationErr(validationErr)
		}

		if unauthorizedErr, ok := err.(*core.ErrUnauthorized); ok {
			return nil, grpc.NewUnauthorizedErr(unauthorizedErr)
		}

		return nil, grpc.NewInternalErr(err)
	}

	return nil, nil
}
