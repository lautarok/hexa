package query

import (
	"context"

	"github.com/lautarok/hexa/src/core/domain"
	"github.com/lautarok/hexa/src/core/errors"
)

type GetPermissionsUsecase struct {
	permissionsRepository domain.IPermissionsRepository
}

type GetPermissionsUsecaseDeps struct {
	PermissionsRepository domain.IPermissionsRepository
}

func NewGetPermissionsUsecase(deps *GetPermissionsUsecaseDeps) *GetPermissionsUsecase {
	return &GetPermissionsUsecase{
		permissionsRepository: deps.PermissionsRepository,
	}
}

type GetPermissionsUsecaseInput struct {
	Page  int
	Limit int
}

func (usecase *GetPermissionsUsecase) GetPermissions(ctx context.Context, input *GetPermissionsUsecaseInput) ([]*domain.Permission, *domain.AppError) {
	permissions, err := usecase.permissionsRepository.FindMany(
		ctx,
		input.Limit*(input.Page-1),
		input.Limit,
	)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return permissions, nil
}
