package account_service

import (
	"context"

	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc"
	"github.com/charmingruby/swrc/proto/pb"
)

func (h *AccountGRPCService) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	input := dto.AuthenticateInputDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.accountService.AuthenticateUseCase(input)
	if err != nil {
		invalidCredentialsErr, ok := err.(*core.ErrInvalidCredentials)
		if ok {
			return nil, grpc.NewInvalidCredentials(invalidCredentialsErr)
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
		return nil, grpc.NewInternalErr(err)
	}

	rep := pb.AuthenticateReply{AccessToken: accessToken}

	return &rep, nil
}
