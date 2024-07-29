package core

import "fmt"

func NewInternalErr(location, reason string) error {
	return &ErrInternal{
		Message: fmt.Sprintf("internal error on %s: %s", location, reason),
	}
}

type ErrInternal struct {
	Message string `json:"message"`
}

func (e *ErrInternal) Error() string {
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
