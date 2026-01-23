package command

import (
	"context"

	"github.com/lautarok/hexa/src/application/domain"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
)

type LoginUsecase struct {
	credentialsRepository domain.ICredentialsRepository
	identityAdapter       ports.IdentityPort
}

type LoginUsecaseDeps struct {
	CredentialsRepository domain.ICredentialsRepository
	IdentityAdapter       ports.IdentityPort
}

func NewLoginUsecase(deps *LoginUsecaseDeps) *LoginUsecase {
	return &LoginUsecase{
		credentialsRepository: deps.CredentialsRepository,
		identityAdapter:       deps.IdentityAdapter,
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
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	if matchCredential == nil {
		return nil, errors.NewNotFoundError("User not found")
	}

	token, exp, err := usecase.identityAdapter.NewToken(&domain.Identity{
		SubUserID: matchCredential.User.ID,
		UserID:    matchCredential.User.ID,
		Email:     matchCredential.Email,
		Username:  matchCredential.Username,
	})

	return &LoginUsecaseOutput{
		Token: token,
		Exp:   exp,
		User:  matchCredential.User,
	}, nil
}
