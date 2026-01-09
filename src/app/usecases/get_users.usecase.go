package usecases

import (
	"context"
	"log"

	"github.com/lautarok/hexa/src/app/domain"
	dto "github.com/lautarok/hexa/src/app/dtos"
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

func (service *GetUsersUsecase) GetUserList(
	ctx context.Context,
	paginationInput *dto.PaginationInput,
) (*dto.UserListOutputDto, error) {
	users, err := service.usersRepository.FindMany(
		ctx,
		paginationInput.Limit*(paginationInput.Page-1),
		paginationInput.Limit,
	)
	if err != nil {
		log.Fatal(err)
	}

	usersDto := []*dto.UserOutputDto{}

	for _, user := range users {
		usersDto = append(usersDto, &dto.UserOutputDto{
			ID:      user.ID,
			Name:    user.Name,
			Surname: user.Surname,
			Credential: &dto.CredentialOutputDto{
				ID:       user.Credential.ID,
				Email:    user.Credential.Email,
				Username: user.Credential.Username,
			},
			CreatedAt: user.CreatedAt,
		})
	}

	output := dto.UserListOutputDto{
		Page:  paginationInput.Page,
		Limit: paginationInput.Limit,
		Users: usersDto,
	}

	return &output, nil
}
