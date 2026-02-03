package errors

import "github.com/lautarok/hexa/src/domain/models"

var ErrorAlreadyExists = &models.AppError{
	Code:    "AlreadyExists",
	Message: "The resource already exists",
}

func NewAlreadyExistsError(message string) *models.AppError {
	return &models.AppError{
		Code:    ErrorAlreadyExists.Code,
		Message: message,
	}
}
