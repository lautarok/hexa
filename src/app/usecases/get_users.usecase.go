package usecases

import (
	"context"
	"log"

	"github.com/lautarok/hexa/src/app/domain"
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

func (service *GetUsersUsecase) GetUserList(ctx context.Context) ([]*domain.User, error) {
	users, err := service.usersRepository.FindMany(ctx, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	return users, nil
}
