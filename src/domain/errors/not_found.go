package errors

import "github.com/lautarok/hexa/src/domain/models"

var ErrorNotFound = &models.AppError{
	Code:    "ResourceNotFound",
	Message: "The resource not found",
}

func NewNotFoundError(message string) *models.AppError {
	return &models.AppError{
		Code:    ErrorNotFound.Code,
		Message: message,
	}
}
