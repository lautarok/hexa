package errors

import "github.com/lautarok/hexa/src/domain/models"

var ErrorInvalidInput = &models.AppError{
	Code:    "InvalidInput",
	Message: "Invalid input",
}

func NewInvalidInputError(message string) *models.AppError {
	return &models.AppError{
		Code:    ErrorInvalidInput.Code,
		Message: message,
	}
}
