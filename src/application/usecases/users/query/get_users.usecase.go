package query

import (
	"context"

	"github.com/lautarok/hexa/src/application/domain"
	"github.com/lautarok/hexa/src/application/errors"
)

type GetUsersUsecase struct {
	usersRepository domain.IUsersRepository
}

type GetUsersUsecaseDeps struct {
	UsersRepository domain.IUsersRepository
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
) ([]*domain.User, *domain.AppError) {
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
