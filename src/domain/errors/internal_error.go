package errors

import "github.com/lautarok/hexa/src/domain/models"

var ErrorInternal = &models.AppError{
	Code:    "InternalError",
	Message: "Application internal error",
}

func NewInternalError(err error) *models.AppError {
	return &models.AppError{
		Code:    ErrorInternal.Code,
		Message: err.Error(),
	}
}
