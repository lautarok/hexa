package ports

import (
	"context"
)

type GoogleOAuth2HandleCallbackResult struct {
	GoogleID   string
	FamilyName string
	GivenName  string
	Email      string
}

type GoogleOAuth2Port interface {
	GetLoginURL(ctx context.Context, locale string) string
	HandleCallback(ctx context.Context, locale string, code string) (*GoogleOAuth2HandleCallbackResult, error)
	IsInvalidGrantError(err error) bool
}
