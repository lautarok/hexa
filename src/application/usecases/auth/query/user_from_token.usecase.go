package query

import (
	"context"

	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/lautarok/hexa/src/domain/ports"
)

type GetUserFromTokenUsecase struct {
	usersRepository       models.IUsersRepository
	credentialsRepository models.ICredentialsRepository
	IdentityAdapter       ports.IdentityPort
}

type GetUserFromTokenUsecaseDeps struct {
	UsersRepository       models.IUsersRepository
	CredentialsRepository models.ICredentialsRepository
	IdentityAdapter       ports.IdentityPort
}

func NewGetUserFromTokenUsecase(deps *GetUserFromTokenUsecaseDeps) *GetUserFromTokenUsecase {
	return &GetUserFromTokenUsecase{
		usersRepository:       deps.UsersRepository,
		credentialsRepository: deps.CredentialsRepository,
		IdentityAdapter:       deps.IdentityAdapter,
	}
}

type GetUserFromTokenUsecaseInput struct {
	Token string
}

type GetUserFromTokenOutput struct {
	User *models.User
}

func (usecase *GetUserFromTokenUsecase) GetUserFromToken(ctx context.Context, input *GetUserFromTokenUsecaseInput) (*GetUserFromTokenOutput, *models.AppError) {
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
