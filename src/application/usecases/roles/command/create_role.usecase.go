package command

import (
	"context"
	"errors"

	"github.com/lautarok/hexa/src/application/domain"
	appErrors "github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
)

type CreateRoleUsecase struct {
	rolesRepository       domain.IRolesRepository
	permissionsRepository domain.IPermissionsRepository
	persistenceAdapter    ports.PersistencePort
}

type CreateRoleUsecaseDeps struct {
	RolesRepository       domain.IRolesRepository
	PermissionsRepository domain.IPermissionsRepository
	PersistenceAdapter    ports.PersistencePort
}

func NewCreateRoleUsecase(deps *CreateRoleUsecaseDeps) *CreateRoleUsecase {
	return &CreateRoleUsecase{
		rolesRepository:       deps.RolesRepository,
		permissionsRepository: deps.PermissionsRepository,
		persistenceAdapter:    deps.PersistenceAdapter,
	}
}

type CreateRoleUsecaseInputPermission struct {
	Alias string
}

type CreateRoleUsecaseInput struct {
	NameEn      string
	NameEs      string
	NameFr      string
	NamePt      string
	Permissions []*CreateRoleUsecaseInputPermission
}

func (usecase *CreateRoleUsecase) CreateRole(ctx context.Context, input *CreateRoleUsecaseInput) (*domain.Role, *domain.AppError) {
	var role *domain.Role
	var repoErr error
	err := usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		permissionAliases := []string{}
		for _, permission := range input.Permissions {
			permissionAliases = append(permissionAliases, permission.Alias)
		}
		permissions, err := usecase.permissionsRepository.FindManyByAlias(ctx, permissionAliases...)
		if err != nil {
			return err
		}

		if len(permissions) == 0 {
			return errors.New("permissions not found")
		}

		role, repoErr = usecase.rolesRepository.CreateOne(ctx, &domain.Role{
			NameEs:      input.NameEs,
			NameEn:      input.NameEn,
			NameFr:      input.NameFr,
			NamePt:      input.NamePt,
			Permissions: permissions,
		})
		if repoErr != nil {
			return repoErr
		}

		return nil
	})

	if err != nil {
		if err.Error() == "permissions not found" {
			return nil, appErrors.NewNotFoundError("Permissions not found")
		} else if usecase.persistenceAdapter.IsUniqueViolation(err) {
			return nil, appErrors.NewAlreadyExistsError("Role already exists")
		} else {
			return nil, appErrors.NewInternalError(err)
		}
	}

	return role, nil
}
