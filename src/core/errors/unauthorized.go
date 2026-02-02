package errors

import "github.com/lautarok/hexa/src/core/domain"

var ErrorUnauthorized = &domain.AppError{
	Code:    "Unauthorized",
	Message: "unauthorized",
}

func NewUnauthorizedError(message string) *domain.AppError {
	return &domain.AppError{
		Code:    ErrorUnauthorized.Code,
		Message: message,
	}
}
