package account_service

import (
	"context"

	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/proto/pb"
)

func (h *AccountGRPCService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	input := dto.RegisterInputDTO{
		GithubDisplayName: req.GithubDisplayName,
		Email:             req.Email,
		Password:          req.Password,
	}

	output, err := h.accountService.RegisterUseCase(input)

	if err != nil {
		if conflictErr, ok := err.(*core.ErrConflict); ok {
			return nil, grpc.NewConflictErr(conflictErr)
		}

		if validationErr, ok := err.(*core.ErrValidation); ok {
			return nil, grpc.NewValidationErr(validationErr)
		}

		return nil, grpc.NewInternalErr(err)
	}

	accessToken, err := h.tokenService.GenerateToken(auth.TokenPayload{
		AccountID: output.ID,
		Role:      output.Role,
		IsValid:   output.IsValid,
		Verified:  output.Verified,
	})
	if err != nil {
		return nil, grpc.NewValidationErr(err)
	}

	rep := pb.RegisterReply{AccessToken: accessToken}

	return &rep, nil
}
