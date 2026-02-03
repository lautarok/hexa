package query

import (
	"context"

	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
)

type GetRolesUsecase struct {
	rolesRepository models.IRolesRepository
}

type GetRolesUsecaseDeps struct {
	RolesRepository models.IRolesRepository
}

func NewGetRolesUsecase(deps *GetRolesUsecaseDeps) *GetRolesUsecase {
	return &GetRolesUsecase{
		rolesRepository: deps.RolesRepository,
	}
}

type GetRolesUsecaseInput struct {
	Page  int
	Limit int
}

func (usecase *GetRolesUsecase) GetRoleList(ctx context.Context, input *GetRolesUsecaseInput) ([]*models.Role, error) {
	roleList, err := usecase.rolesRepository.FindMany(
		ctx,
		input.Limit*(input.Page-1),
		input.Limit,
	)

	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return roleList, nil
}
