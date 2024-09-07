package account_rpc

import (
	"context"

	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/interceptor"
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
