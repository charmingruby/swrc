package core

import "fmt"

func NewInternalErr() error {
	return &ErrInternal{
		Message: "internal error",
	}
}

type ErrInternal struct {
	Message string `json:"message"`
}

func (e *ErrInternal) Error() string {
	return e.Message
}

func NewAlreadyExistsErr(entity string) error {
	return &ErrAlreadyExists{
		Message: fmt.Sprintf("%s already exists error", entity),
	}
}

type ErrAlreadyExists struct {
	Message string `json:"message"`
}

func (e *ErrAlreadyExists) Error() string {
	return e.Message
}

func NewConflictErr(entity, field string) error {
	return &ErrConflict{
		Message: fmt.Sprintf("%s %s is already taken", entity, field),
	}
}

type ErrConflict struct {
	Message string `json:"message"`
}

func (e *ErrConflict) Error() string {
	return e.Message
}

func NewNotFoundErr(entity string) error {
	return &ErrNotFound{
		Message: fmt.Sprintf("%s not found", entity),
	}
}

type ErrNotFound struct {
	Message string `json:"message"`
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

func NewInvalidCredentialsErr() error {
	return &ErrInvalidCredentials{
		Message: "invalid credentials",
	}
}

type ErrInvalidCredentials struct {
	Message string `json:"message"`
}

func (e *ErrInvalidCredentials) Error() string {
	return e.Message
}

func NewUnauthorizedErr() error {
	return &ErrUnauthorized{
		Message: "unauthorized",
	}
}

type ErrUnauthorized struct {
	Message string `json:"message"`
}

func (e *ErrUnauthorized) Error() string {
	return e.Message
}
