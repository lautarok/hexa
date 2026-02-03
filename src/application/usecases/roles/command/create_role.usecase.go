package command

import (
	"context"
	"errors"

	appErrors "github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/lautarok/hexa/src/domain/ports"
)

type CreateRoleUsecase struct {
	rolesRepository       models.IRolesRepository
	permissionsRepository models.IPermissionsRepository
	persistenceAdapter    ports.PersistencePort
}

type CreateRoleUsecaseDeps struct {
	RolesRepository       models.IRolesRepository
	PermissionsRepository models.IPermissionsRepository
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

func (usecase *CreateRoleUsecase) CreateRole(ctx context.Context, input *CreateRoleUsecaseInput) (*models.Role, *models.AppError) {
	var role *models.Role
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

		rolePermissions := []models.Permission{}
		for _, permission := range permissions {
			rolePermissions = append(rolePermissions, *permission)
		}

		role, repoErr = usecase.rolesRepository.CreateOne(ctx, &models.Role{
			NameEs:      input.NameEs,
			NameEn:      input.NameEn,
			NameFr:      input.NameFr,
			NamePt:      input.NamePt,
			Permissions: rolePermissions,
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
