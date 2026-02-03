package query

import (
	"context"

	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
)

type GetPermissionsUsecase struct {
	permissionsRepository models.IPermissionsRepository
}

type GetPermissionsUsecaseDeps struct {
	PermissionsRepository models.IPermissionsRepository
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

func (usecase *GetPermissionsUsecase) GetPermissions(ctx context.Context, input *GetPermissionsUsecaseInput) ([]*models.Permission, *models.AppError) {
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
