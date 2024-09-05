package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewInternalErr(err error) error {
	return status.Errorf(codes.Internal, err.Error())
}

func NewUnauthorizedErr(err error) error {
	return status.Errorf(codes.Unauthenticated, err.Error())
}

func NewValidationErr(err error) error {
	return status.Errorf(codes.InvalidArgument, err.Error())
}

func NewNotFoundErr(err error) error {
	return status.Errorf(codes.NotFound, err.Error())
}

func NewConflictErr(err error) error {
	return status.Errorf(codes.AlreadyExists, err.Error())
}

func NewInvalidCredentials(err error) error {
	return status.Errorf(codes.InvalidArgument, err.Error())
}
