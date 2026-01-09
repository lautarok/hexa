package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/app/domain"
	dto "github.com/lautarok/hexa/src/app/dtos"
)

type ErrorMiddleware struct{}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

func sendInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, &dto.AppErrorDto{
		StatusCode: 500,
		Code:       "InternalError",
		Message:    "Internal server error",
	})
}

func (middleware *ErrorMiddleware) HandleErrors(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) == 0 {
		return
	}

	lastError := ctx.Errors.Last().Err
	if lastError != nil {
		appError, ok := lastError.(*domain.AppError)

		if !ok {
			sendInternalError(ctx)
			return
		}

		statusCode := 500
		switch appError.Code {
		case "NotFound":
			statusCode = 404
		case "InvalidInput":
			statusCode = 400
		case "AlreadyExists":
			statusCode = 409
		}

		ctx.JSON(statusCode, &dto.AppErrorDto{
			StatusCode: statusCode,
			Code:       appError.Code,
			Message:    appError.Message,
		})
	}
}
