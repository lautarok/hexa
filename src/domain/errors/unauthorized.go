package errors

import "github.com/lautarok/hexa/src/domain/models"

var ErrorUnauthorized = &models.AppError{
	Code:    "Unauthorized",
	Message: "unauthorized",
}

func NewUnauthorizedError(message string) *models.AppError {
	return &models.AppError{
		Code:    ErrorUnauthorized.Code,
		Message: message,
	}
}
