package ports

import (
	"context"

	"github.com/lautarok/hexa/src/core/ports/dtos"
)

type GoogleOAuth2Port interface {
	GetLoginURL(ctx context.Context) string
	HandleCallback(ctx context.Context, code string) (*dtos.GoogleIdentityDto, error)
	IsInvalidGrantError(err error) bool
}
