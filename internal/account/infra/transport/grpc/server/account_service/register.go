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

func (h *AccountGRPCService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	input := dto.RegisterInputDTO{
		GithubDisplayName: req.GithubDisplayName,
		Email:             req.Email,
		Password:          req.Password,
	}

	output, err := h.accountService.RegisterUseCase(input)

	if err != nil {
		conflictErr, ok := err.(*core.ErrConflict)
		if ok {
			return nil, status.Errorf(codes.AlreadyExists, conflictErr.Error())
		}

		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			return nil, status.Errorf(codes.InvalidArgument, validationErr.Error())
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

	rep := pb.RegisterReply{AccessToken: accessToken}

	return &rep, nil
}
