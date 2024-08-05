package account_service

import (
	"context"

	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/infra/auth"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			return nil, status.Errorf(codes.InvalidArgument, invalidCredentialsErr.Message)
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	accessToken, err := h.tokenService.GenerateToken(auth.TokenPayload{
		AccountID: output.ID,
		Role:      output.Role,
		IsValid:   output.IsValid,
		Verified:  output.Verified,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	rep := pb.AuthenticateReply{AccessToken: accessToken}

	return &rep, nil
}
