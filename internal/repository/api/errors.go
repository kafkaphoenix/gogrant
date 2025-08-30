package api

import (
	"errors"
	"fmt"
)

var (
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrInvalidPort         = errors.New("invalid port")
)

// ServerError represents http server errors.
type ServerError struct {
	Message string
	Err     error
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *ServerError) Unwrap() error {
	return e.Err
}

// RouterError represents router errors.
type RouterError struct {
}

func (e *RouterError) Error() string {
	return "failed to get router from server"
}

// ValidationError represents validation errors.
type ValidationError struct {
	Param string
	Err   error
}

func (e *ValidationError) Error() string {
	return "validation error: " + e.Param + ": " + e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}
