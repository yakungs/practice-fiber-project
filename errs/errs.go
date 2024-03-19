package errs

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (err AppError) Error() string {
	return err.Message
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnExpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "Unexpected error",
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
