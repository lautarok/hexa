package query

import (
	"context"

	"github.com/lautarok/hexa/src/core/ports"
)

type GetGoogleOAuth2URLUsecase struct {
	googleOauth2Adapter ports.GoogleOAuth2Port
}

type GetGoogleOAuth2URLUsecaseDeps struct {
	GoogleOAuth2Adapter ports.GoogleOAuth2Port
}

func NewGetGoogleOAuth2URLUsecase(deps *GetGoogleOAuth2URLUsecaseDeps) *GetGoogleOAuth2URLUsecase {
	return &GetGoogleOAuth2URLUsecase{
		googleOauth2Adapter: deps.GoogleOAuth2Adapter,
	}
}

func (usecase *GetGoogleOAuth2URLUsecase) GetGoogleAuthURL(ctx context.Context) string {
	return usecase.googleOauth2Adapter.GetLoginURL(ctx)
}
