package errors

import "github.com/lautarok/hexa/src/app/domain"

var ErrorAlreadyExists = &domain.AppError{
	Code:    "AlreadyExists",
	Message: "The resource already exists",
}

func NewAlreadyExistsError(message string) *domain.AppError {
	return &domain.AppError{
		Code:    ErrorAlreadyExists.Code,
		Message: message,
	}
}
