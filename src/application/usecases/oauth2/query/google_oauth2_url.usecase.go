package query

import (
	"context"

	"github.com/lautarok/hexa/src/domain/ports"
)

type GetGoogleSignOnURLUsecase struct {
	googleOauth2Adapter ports.GoogleOAuth2Port
}

type GetGoogleSignOnURLUsecaseDeps struct {
	GoogleOAuth2Adapter ports.GoogleOAuth2Port
}

func NewGetGoogleSignOnURLUsecase(deps *GetGoogleSignOnURLUsecaseDeps) *GetGoogleSignOnURLUsecase {
	return &GetGoogleSignOnURLUsecase{
		googleOauth2Adapter: deps.GoogleOAuth2Adapter,
	}
}

func (usecase *GetGoogleSignOnURLUsecase) GetGoogleAuthURL(ctx context.Context, locale string) string {
	return usecase.googleOauth2Adapter.GetLoginURL(ctx, locale)
}
