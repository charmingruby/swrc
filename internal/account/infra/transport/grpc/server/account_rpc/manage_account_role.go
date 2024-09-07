package account_rpc

import (
	"context"
	"errors"

	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/interceptor"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *AccountGRPCService) ManageAccountRole(ctx context.Context, req *pb.ManageAccountRoleRequest) (*emptypb.Empty, error) {
	managerID := ctx.Value(interceptor.AccountIDKey).(string)
	if managerID == "" {
		return nil, grpc.NewInternalErr(errors.New("manager id not found"))
	}

	input := dto.ManageAccountRoleInputDTO{
		ManagerAccountID: managerID,
		AccountID:        req.AccountId,
		NewRole:          req.NewRole,
	}

	if err := h.accountService.ManageAccountRoleUseCase(input); err != nil {
		if notFoundErr, ok := err.(*core.ErrNotFound); ok {
			return nil, grpc.NewNotFoundErr(notFoundErr)
		}

		if unauthorizedErr, ok := err.(*core.ErrUnauthorized); ok {
			return nil, grpc.NewUnauthorizedErr(unauthorizedErr)
		}

		return nil, grpc.NewInternalErr(err)
	}

	return nil, nil
}
