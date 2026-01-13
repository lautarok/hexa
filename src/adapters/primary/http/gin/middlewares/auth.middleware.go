package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/usecases"
)

type AuthMiddleware struct {
	getUserFromTokenUsecase *usecases.GetUserFromTokenUsecase
}

type AuthMiddlewareDeps struct {
	GetUserFromTokenUsecase *usecases.GetUserFromTokenUsecase
}

func NewAuthMiddleware(deps *AuthMiddlewareDeps) *AuthMiddleware {
	return &AuthMiddleware{
		getUserFromTokenUsecase: deps.GetUserFromTokenUsecase,
	}
}

func (middleware *AuthMiddleware) HandleAuth(permissionAlias ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		token, havePrefix := strings.CutPrefix(authHeader, "Bearer ")

		if !havePrefix {
			ctx.Error(errors.NewUnauthorizedError("Invalid token"))
			ctx.Abort()
			return
		}

		userFromToken, appError := middleware.getUserFromTokenUsecase.GetUserFromToken(ctx, &usecases.GetUserFromTokenInput{
			Token: token,
		})

		if appError != nil {
			ctx.Error(appError)
			ctx.Abort()
			return
		}

		for _, requiredPermission := range permissionAlias {
			haveThisPermission := false
			for _, havePermission := range userFromToken.User.Role.Permissions {
				if requiredPermission == havePermission.Alias {
					haveThisPermission = true
					break
				}
			}

			if haveThisPermission == false {
				ctx.Error(errors.NewUnauthorizedError("Insufficient permissions"))
				ctx.Abort()
				return
			}
		}

		ctx.Set("auth_user", userFromToken.User)
	}
}
