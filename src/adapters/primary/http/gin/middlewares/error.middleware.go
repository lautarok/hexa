package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/adapters/primary/http/gin/dtos"
	"github.com/lautarok/hexa/src/application/domain"
)

type ErrorMiddleware struct{}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

func sendInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, &dtos.AppErrorDto{
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
		case "ResourceNotFound":
			statusCode = 404
		case "InvalidInput":
			statusCode = 400
		case "AlreadyExists":
			statusCode = 409
		case "Unauthorized":
			statusCode = 401
		case "InternalError":
			log.Printf(
				"%s [%s] INTERNAL ERROR: %s",
				ctx.FullPath(),
				ctx.Request.Method,
				appError.Message,
			)
			sendInternalError(ctx)
			return
		}

		ctx.JSON(statusCode, &dtos.AppErrorDto{
			StatusCode: statusCode,
			Code:       appError.Code,
			Message:    appError.Message,
		})
	}
}
