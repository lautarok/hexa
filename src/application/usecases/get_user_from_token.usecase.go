package usecases

import (
	"context"

	"github.com/lautarok/hexa/src/application/domain"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
)

type GetUserFromTokenUsecase struct {
	usersRepository       domain.IUsersRepository
	credentialsRepository domain.ICredentialsRepository
	IdentityAdapter       ports.IdentityPort
}

type GetUserFromTokenUsecaseDeps struct {
	UsersRepository       domain.IUsersRepository
	CredentialsRepository domain.ICredentialsRepository
	IdentityAdapter       ports.IdentityPort
}

func NewGetUserFromTokenUsecase(deps *GetUserFromTokenUsecaseDeps) *GetUserFromTokenUsecase {
	return &GetUserFromTokenUsecase{
		usersRepository:       deps.UsersRepository,
		credentialsRepository: deps.CredentialsRepository,
		IdentityAdapter:       deps.IdentityAdapter,
	}
}

type GetUserFromTokenInput struct {
	Token string
}

type GetUserFromTokenOutput struct {
	User *domain.User
}

func (usecase *GetUserFromTokenUsecase) GetUserFromToken(ctx context.Context, input *GetUserFromTokenInput) (*GetUserFromTokenOutput, *domain.AppError) {
	payload, err := usecase.IdentityAdapter.ParseToken(input.Token)
	if err != nil {
		return nil, errors.NewUnauthorizedError("Invalid token")
	}

	user, err := usecase.usersRepository.FindOneByID(ctx, payload.UserID)
	if err != nil {
		return nil, errors.NewInternalError(err)
	} else if user == nil {
		return nil, errors.NewNotFoundError("User not found")
	}

	return &GetUserFromTokenOutput{
		User: user,
	}, nil
}
