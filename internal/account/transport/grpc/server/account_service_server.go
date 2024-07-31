package server

import (
	"context"

	"github.com/charmingruby/swrc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewGRPCAccountServiceHandler() *GRPCAccountServiceHandler {
	return &GRPCAccountServiceHandler{}
}

type GRPCAccountServiceHandler struct {
	pb.UnimplementedAccountServiceServer
}

func (h *GRPCAccountServiceHandler) Authenticate(context.Context, *pb.AuthenticateRequest) (*pb.AuthenticateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func (h *GRPCAccountServiceHandler) Register(context.Context, *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (h *GRPCAccountServiceHandler) ManageAccountRole(context.Context, *pb.ManageAccountRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageAccountRole not implemented")
}

func (h *GRPCAccountServiceHandler) VerifyAccount(context.Context, *pb.VerifyAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccount not implemented")
}
