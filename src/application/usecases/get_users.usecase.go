package usecases

import (
	"context"
	"log"

	"github.com/lautarok/hexa/src/application/domain"
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

type GetUserListInput struct {
	Page  int
	Limit int
}

// SACAR DTOs
func (service *GetUsersUsecase) GetUserList(
	ctx context.Context,
	input *GetUserListInput,
) ([]*domain.User, error) {
	users, err := service.usersRepository.FindMany(
		ctx,
		input.Limit*(input.Page-1),
		input.Limit,
	)
	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}
