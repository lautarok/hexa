package query

import (
	"context"

	"github.com/lautarok/hexa/src/core/domain"
	"github.com/lautarok/hexa/src/core/errors"
	"github.com/lautarok/hexa/src/core/ports"
	"github.com/lautarok/hexa/src/core/ports/dtos"
)

type GetGoogleOAuth2UserUsecase struct {
	googleOauth2Adapter ports.GoogleOAuth2Port
}

type GetGoogleOAuth2UserUsecaseDeps struct {
	GoogleOAuth2Adapter ports.GoogleOAuth2Port
}

func NewGetGoogleOAuth2UserUsecase(deps *GetGoogleOAuth2UserUsecaseDeps) *GetGoogleOAuth2UserUsecase {
	return &GetGoogleOAuth2UserUsecase{
		googleOauth2Adapter: deps.GoogleOAuth2Adapter,
	}
}

func (usecase *GetGoogleOAuth2UserUsecase) GetGoogleUser(ctx context.Context, code string) (*dtos.GoogleIdentityDto, *domain.AppError) {
	identity, err := usecase.googleOauth2Adapter.HandleCallback(ctx, code)
	if err != nil {
		if usecase.googleOauth2Adapter.IsInvalidGrantError(err) {
			return nil, errors.NewInvalidInputError("invalid code")
		}
		return nil, errors.NewInternalError(err)
	}

	return identity, nil
}
