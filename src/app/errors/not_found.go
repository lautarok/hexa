package errors

import "github.com/lautarok/hexa/src/app/domain"

var ErrorNotFound = &domain.AppError{
	Code:    "ResourceNotFound",
	Message: "The resource not found",
}

func NewNotFoundError(message string) *domain.AppError {
	return &domain.AppError{
		Code:    ErrorNotFound.Code,
		Message: message,
	}
}
