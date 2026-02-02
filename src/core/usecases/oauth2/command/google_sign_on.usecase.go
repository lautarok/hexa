package command

import (
	"context"

	"github.com/lautarok/hexa/src/core/domain"
	"github.com/lautarok/hexa/src/core/errors"
	"github.com/lautarok/hexa/src/core/ports"
)

type GoogleSignOnUsecase struct {
	usersRepository            domain.IUsersRepository
	credentialsRepository      domain.ICredentialsRepository
	googleIdentitiesRepository domain.IGoogleIdentitiesRepository
	persistenceAdapter         ports.PersistencePort
	googleOauth2Adapter        ports.GoogleOAuth2Port
	identityAdapter            ports.IdentityPort
}

type GoogleSignOnUsecaseDeps struct {
	UsersRepository            domain.IUsersRepository
	CredentialsRepository      domain.ICredentialsRepository
	GoogleIdentitiesRepository domain.IGoogleIdentitiesRepository
	PersistenceAdapter         ports.PersistencePort
	GoogleOAuth2Adapter        ports.GoogleOAuth2Port
	IdentityAdapter            ports.IdentityPort
}

func NewGoogleSignOnUsecase(deps *GoogleSignOnUsecaseDeps) *GoogleSignOnUsecase {
	return &GoogleSignOnUsecase{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

type GoogleSignOnUsecaseInput struct {
	code string
}

type GoogleSignOnUsecaseOutput struct {
	Token string
	Exp   int64
	User  *domain.User
}

func (usecase *GoogleSignOnUsecase) SignOn(ctx context.Context, code string) (*GoogleSignOnUsecaseOutput, *domain.AppError) {
	googleIdentity, err := usecase.googleOauth2Adapter.HandleCallback(ctx, code)
	if err != nil {
		if usecase.googleOauth2Adapter.IsInvalidGrantError(err) {
			return nil, errors.NewInvalidInputError("invalid code")
		}
		return nil, errors.NewInternalError(err)
	}

	matchGoogleIdentity, err := usecase.googleIdentitiesRepository.FindOneByGoogleID(ctx, googleIdentity.GoogleID)
	if err == nil {
		token, exp, err := usecase.identityAdapter.NewToken(&domain.Identity{
			SubUserID: matchGoogleIdentity.User.ID,
			UserID:    matchGoogleIdentity.User.ID,
			Email:     matchGoogleIdentity.User.Credential.Email,
			Username:  matchGoogleIdentity.User.Credential.Username,
		})
		if err != nil {
			return nil, errors.NewInternalError(err)
		}

		return &GoogleSignOnUsecaseOutput{
			Token: token,
			Exp:   exp,
			User:  &matchGoogleIdentity.User,
		}, nil
	}

	usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		usecase.googleIdentitiesRepository.CreateOne(ctx, &domain.GoogleIdentity{
			GoogleID: googleIdentity.GoogleID,
			User: &domain.User{
				ID: googleIdentity,
			},
		})
	})

	return nil, nil
}
