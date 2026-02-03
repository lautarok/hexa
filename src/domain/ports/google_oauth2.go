package ports

import (
	"context"

	"github.com/lautarok/hexa/src/domain/ports/dtos"
)

type GoogleOAuth2Port interface {
	GetLoginURL(ctx context.Context, locale string) string
	HandleCallback(ctx context.Context, locale string, code string) (*dtos.GoogleIdentityDto, error)
	IsInvalidGrantError(err error) bool
}
