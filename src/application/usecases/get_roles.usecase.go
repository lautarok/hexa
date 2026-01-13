package usecases

import (
	"context"
	"log"

	"github.com/lautarok/hexa/src/application/domain"
	"github.com/lautarok/hexa/src/application/errors"
)

type GetRolesUsecase struct {
	rolesRepository domain.IRolesRepository
}

type GetRolesUsecaseDeps struct {
	RolesRepository domain.IRolesRepository
}

func NewGetRolesUsecase(deps *GetRolesUsecaseDeps) *GetRolesUsecase {
	return &GetRolesUsecase{
		rolesRepository: deps.RolesRepository,
	}
}

type GetRolesInput struct {
	Page  int
	Limit int
}

func (usecase *GetRolesUsecase) GetRoleList(ctx context.Context, input *GetRolesInput) ([]*domain.Role, error) {
	roleList, err := usecase.rolesRepository.FindMany(
		ctx,
		input.Limit*(input.Page-1),
		input.Limit,
	)

	if err != nil {
		log.Println("111111111111111")
		log.Println(input)
		log.Println(usecase.rolesRepository)
		log.Println("111111111111111")
		return nil, errors.NewInternalError(err)
	}

	return roleList, nil
}
