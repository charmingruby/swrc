package server

import (
	"context"

	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/usecase"
	"github.com/charmingruby/swrc/internal/account/infra/security"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *AccountGRPCServerHandler) newAccountServiceGRPCServerHandler() *AccountServiceGRPCServerHandler {
	return &AccountServiceGRPCServerHandler{
		accountService: h.accountService,
		tokenService:   h.tokenService,
	}
}

type AccountServiceGRPCServerHandler struct {
	pb.UnimplementedAccountServiceServer

	accountService usecase.AccountUseCase
	tokenService   security.TokenService
}

func (h *AccountServiceGRPCServerHandler) Authenticate(context.Context, *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func (h *AccountServiceGRPCServerHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
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

	accessToken, err := h.tokenService.GenerateToken(security.TokenPayload{
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

func (h *AccountServiceGRPCServerHandler) ManageAccountRole(context.Context, *pb.ManageAccountRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageAccountRole not implemented")
}

func (h *AccountServiceGRPCServerHandler) VerifyAccount(context.Context, *pb.VerifyAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccount not implemented")
}
