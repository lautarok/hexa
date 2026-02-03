package command

import (
	"context"

	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/lautarok/hexa/src/domain/ports"
	"github.com/lautarok/hexa/src/pkg"
)

type GoogleSignOnUsecase struct {
	usersRepository            models.IUsersRepository
	rolesRepository            models.IRolesRepository
	credentialsRepository      models.ICredentialsRepository
	googleIdentitiesRepository models.IGoogleIdentitiesRepository
	persistenceAdapter         ports.PersistencePort
	googleOauth2Adapter        ports.GoogleOAuth2Port
	identityAdapter            ports.IdentityPort
}

type GoogleSignOnUsecaseDeps struct {
	UsersRepository            models.IUsersRepository
	RolesRepository            models.IRolesRepository
	CredentialsRepository      models.ICredentialsRepository
	GoogleIdentitiesRepository models.IGoogleIdentitiesRepository
	PersistenceAdapter         ports.PersistencePort
	GoogleOAuth2Adapter        ports.GoogleOAuth2Port
	IdentityAdapter            ports.IdentityPort
}

func NewGoogleSignOnUsecase(deps *GoogleSignOnUsecaseDeps) *GoogleSignOnUsecase {
	return &GoogleSignOnUsecase{
		usersRepository:            deps.UsersRepository,
		rolesRepository:            deps.RolesRepository,
		credentialsRepository:      deps.CredentialsRepository,
		googleIdentitiesRepository: deps.GoogleIdentitiesRepository,
		persistenceAdapter:         deps.PersistenceAdapter,
		googleOauth2Adapter:        deps.GoogleOAuth2Adapter,
		identityAdapter:            deps.IdentityAdapter,
	}
}

type GoogleSignOnUsecaseInput struct {
	code string
}

type GoogleSignOnUsecaseOutput struct {
	Token string
	Exp   int64
	User  *models.User
}

func (usecase *GoogleSignOnUsecase) GoogleSignOn(ctx context.Context, locale string, code string) (*GoogleSignOnUsecaseOutput, *models.AppError) {
	googleIdentityCallback, err := usecase.googleOauth2Adapter.HandleCallback(ctx, locale, code)
	if err != nil {
		if usecase.googleOauth2Adapter.IsInvalidGrantError(err) {
			return nil, errors.NewInvalidInputError("invalid code")
		}
		return nil, errors.NewInternalError(err)
	}

	matchCredential, err := usecase.credentialsRepository.FindByUsernameOrEmail(ctx, googleIdentityCallback.Email)
	if err == nil {
		_, err := usecase.googleIdentitiesRepository.FindOneByGoogleID(ctx, googleIdentityCallback.GoogleID)
		if err != nil {
			if usecase.persistenceAdapter.IsErrNotFound(err) {
				_, err := usecase.googleIdentitiesRepository.CreateOne(ctx, &models.GoogleIdentity{
					GoogleID: googleIdentityCallback.GoogleID,
					UserID:   matchCredential.UserID,
				})
				if err != nil {
					return nil, errors.NewInternalError(err)
				}
			} else {
				return nil, errors.NewInternalError(err)
			}
		}

		token, exp, err := usecase.identityAdapter.NewToken(&models.Identity{
			SubUserID: matchCredential.UserID,
			UserID:    matchCredential.UserID,
		})
		if err != nil {
			return nil, errors.NewInternalError(err)
		}

		user, err := usecase.usersRepository.FindOneByID(ctx, matchCredential.UserID)
		if err != nil {
			return nil, errors.NewInternalError(err)
		}

		return &GoogleSignOnUsecaseOutput{
			Token: token,
			Exp:   exp,
			User:  user,
		}, nil
	}

	var user *models.User

	err = usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		role, err := usecase.rolesRepository.GetOneBySlug(ctx, "common:user")
		if err != nil {
			return err
		}

		surname := googleIdentityCallback.FamilyName
		if len(surname) > 0 {
			surname = pkg.NormalizeName(surname)
		}

		user, err = usecase.usersRepository.CreateOne(ctx, &models.User{
			Name:    pkg.NormalizeName(googleIdentityCallback.GivenName),
			Surname: surname,
			Role: models.Role{
				ID: role.ID,
			},
		})
		if err != nil {
			return err
		}

		_, err = usecase.credentialsRepository.CreateOne(ctx, &models.Credential{
			Email:  googleIdentityCallback.Email,
			UserID: user.ID,
		})
		if err != nil {
			return err
		}

		_, err = usecase.googleIdentitiesRepository.CreateOne(ctx, &models.GoogleIdentity{
			GoogleID: googleIdentityCallback.GoogleID,
			UserID:   user.ID,
		})
		return err
	})
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	user, err = usecase.usersRepository.FindOneByID(ctx, user.ID)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	token, exp, err := usecase.identityAdapter.NewToken(&models.Identity{
		SubUserID: user.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return &GoogleSignOnUsecaseOutput{
		Token: token,
		Exp:   exp,
		User:  user,
	}, nil
}
