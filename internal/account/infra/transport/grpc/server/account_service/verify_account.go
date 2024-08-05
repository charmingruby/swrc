package account_service

import (
	"context"

	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/infra/auth/interceptor"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *AccountGRPCService) VerifyAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*emptypb.Empty, error) {
	managerID := ctx.Value(interceptor.AccountIDKey).(string)
	if managerID == "" {
		return nil, status.Errorf(codes.Internal, "manager id not found")
	}

	input := dto.VerifyAccountInputDTO{
		SolicitorAccountID: managerID,
		AccountToVerifyID:  req.AccountToVerifyId,
		Verification:       req.Verification,
	}

	if err := h.accountService.VerifyAccountUseCase(input); err != nil {
		notFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			return nil, status.Errorf(codes.NotFound, notFoundErr.Error())
		}

		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			return nil, status.Errorf(codes.InvalidArgument, validationErr.Error())
		}

		unauthorizedErr, ok := err.(*core.ErrUnauthorized)
		if ok {
			return nil, status.Errorf(codes.Unauthenticated, unauthorizedErr.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return nil, nil
}
