package command

import (
	"context"

	"github.com/lautarok/hexa/src/core/domain"
	"github.com/lautarok/hexa/src/core/errors"
	"github.com/lautarok/hexa/src/core/ports"
)

type LoginUsecase struct {
	credentialsRepository domain.ICredentialsRepository
	usersRepository       domain.IUsersRepository
	persistenceAdapter    ports.PersistencePort
	identityAdapter       ports.IdentityPort
	passwordAdapter       ports.PasswordPort
}

type LoginUsecaseDeps struct {
	CredentialsRepository domain.ICredentialsRepository
	UsersRepository       domain.IUsersRepository
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
	User  *domain.User
}

func (usecase *LoginUsecase) Login(ctx context.Context, input *LoginUsecaseInput) (*LoginUsecaseOutput, *domain.AppError) {
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

	token, exp, err := usecase.identityAdapter.NewToken(&domain.Identity{
		SubUserID: matchUser.ID,
		UserID:    matchUser.ID,
		Email:     matchCredential.Email,
		Username:  matchCredential.Username,
	})

	return &LoginUsecaseOutput{
		Token: token,
		Exp:   exp,
		User:  matchUser,
	}, nil
}
