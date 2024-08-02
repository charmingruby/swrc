package server

import (
	"context"

	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newAccountServiceGRPCServerHandler() *AccountServiceGRPCServerHandler {
	return &AccountServiceGRPCServerHandler{}
}

type AccountServiceGRPCServerHandler struct {
	pb.UnimplementedAccountServiceServer
}

func (h *AccountServiceGRPCServerHandler) Authenticate(context.Context, *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func (h *AccountServiceGRPCServerHandler) Register(context.Context, *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (h *AccountServiceGRPCServerHandler) ManageAccountRole(context.Context, *pb.ManageAccountRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageAccountRole not implemented")
}

func (h *AccountServiceGRPCServerHandler) VerifyAccount(context.Context, *pb.VerifyAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccount not implemented")
}
