package command

import (
	"context"

	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/lautarok/hexa/src/domain/ports"
)

type LoginUsecase struct {
	credentialsRepository models.ICredentialsRepository
	usersRepository       models.IUsersRepository
	persistenceAdapter    ports.PersistencePort
	identityAdapter       ports.IdentityPort
	passwordAdapter       ports.PasswordPort
}

type LoginUsecaseDeps struct {
	CredentialsRepository models.ICredentialsRepository
	UsersRepository       models.IUsersRepository
	PersistenceAdapter    ports.PersistencePort
	IdentityAdapter       ports.IdentityPort
	PasswordAdapter       ports.PasswordPort
}

func NewLoginUsecase(deps *LoginUsecaseDeps) *LoginUsecase {
	return &LoginUsecase{
		credentialsRepository: deps.CredentialsRepository,
		usersRepository:       deps.UsersRepository,
		persistenceAdapter:    deps.PersistenceAdapter,
		identityAdapter:       deps.IdentityAdapter,
		passwordAdapter:       deps.PasswordAdapter,
	}
}

type LoginUsecaseInput struct {
	UsernameOrEmail string
	Password        string
}

type LoginUsecaseOutput struct {
	Token string
	Exp   int64
	User  *models.User
}

func (usecase *LoginUsecase) Login(ctx context.Context, input *LoginUsecaseInput) (*LoginUsecaseOutput, *models.AppError) {
	matchCredential, err := usecase.credentialsRepository.FindByUsernameOrEmail(ctx, input.UsernameOrEmail)
	if err != nil && !usecase.persistenceAdapter.IsErrNotFound(err) {
		return nil, errors.NewInternalError(err)
	}

	if matchCredential == nil {
		return nil, errors.NewNotFoundError("User not found")
	}

	if !usecase.passwordAdapter.Compare(matchCredential.Password, input.Password) {
		return nil, errors.NewNotFoundError("User not found")
	}

	matchUser, err := usecase.usersRepository.FindOneByID(ctx, matchCredential.UserID)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	token, exp, err := usecase.identityAdapter.NewToken(&models.Identity{
		SubUserID: matchUser.ID,
		UserID:    matchUser.ID,
	})

	return &LoginUsecaseOutput{
		Token: token,
		Exp:   exp,
		User:  matchUser,
	}, nil
}
