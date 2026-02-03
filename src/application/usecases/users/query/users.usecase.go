package query

import (
	"context"

	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
)

type GetUsersUsecase struct {
	usersRepository models.IUsersRepository
}

type GetUsersUsecaseDeps struct {
	UsersRepository models.IUsersRepository
}

func NewGetUsersUsecase(deps *GetUsersUsecaseDeps) *GetUsersUsecase {
	return &GetUsersUsecase{
		usersRepository: deps.UsersRepository,
	}
}

type GetUsersUsecaseInput struct {
	Page  int
	Limit int
}

func (service *GetUsersUsecase) GetUserList(
	ctx context.Context,
	input *GetUsersUsecaseInput,
) ([]*models.User, *models.AppError) {
	users, err := service.usersRepository.FindMany(
		ctx,
		input.Limit*(input.Page-1),
		input.Limit,
	)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return users, nil
}
