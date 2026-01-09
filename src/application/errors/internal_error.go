package errors

import "github.com/lautarok/hexa/src/application/domain"

var ErrorInternal = &domain.AppError{
	Code:    "InternalError",
	Message: "Application internal error",
}

func NewInternalError(err error) *domain.AppError {
	return &domain.AppError{
		Code:    ErrorInternal.Code,
		Message: err.Error(),
	}
}
