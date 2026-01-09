package errors

import "github.com/lautarok/hexa/src/app/domain"

var ErrorInvalidInput = &domain.AppError{
	Code:    "InvalidInput",
	Message: "Invalid input",
}

func NewInvalidInputError(message string) *domain.AppError {
	return &domain.AppError{
		Code:    ErrorInvalidInput.Code,
		Message: message,
	}
}
