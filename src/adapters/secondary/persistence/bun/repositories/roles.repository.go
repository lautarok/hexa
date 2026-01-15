package repositories

import (
	"context"
	"database/sql"

	bunPersistence "github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/application/domain"
	"github.com/uptrace/bun"
)

type RolesRepository struct {
	dbAdapter *bunPersistence.BunAdapter
}

type RolesRepositoryDeps struct {
	DBAdapter *bunPersistence.BunAdapter
}

func NewRolesRepository(deps *RolesRepositoryDeps) domain.IRolesRepository {
	return &RolesRepository{
		dbAdapter: deps.DBAdapter,
	}
}

func (repository *RolesRepository) GetOneByAlias(ctx context.Context, alias string) (*domain.Role, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var role entities.Role
	err := db.
		NewSelect().
		Model(&role).
		Where("alias = ?", alias).
		Scan(ctx)

	return role.ToDomain(), err
}

func (repository *RolesRepository) FindMany(ctx context.Context, skip int, limit int) ([]*domain.Role, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var entityRoleList []*entities.Role
	err := db.NewSelect().
		Model(&entityRoleList).
		Limit(limit).
		Offset(skip).
		Relation("Permissions").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	roleList := []*domain.Role{}

	for _, entityRole := range entityRoleList {
		roleList = append(roleList, entityRole.ToDomain())
	}

	return roleList, err
}

func (repository *RolesRepository) CreateOne(ctx context.Context, domainRole *domain.Role) (*domain.Role, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var entityRole entities.Role
	entityRole.FromDomain(domainRole)

	rolePermissions := []*entities.RolePermission{}

	err := db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().
			Model(&entityRole).
			Returning("*").
			Exec(ctx)
		if err != nil {
			return err
		}

		for _, permission := range domainRole.Permissions {
			rolePermissions = append(rolePermissions, &entities.RolePermission{
				RoleID:       entityRole.ID,
				PermissionID: permission.ID,
			})
		}

		_, err = tx.NewInsert().
			Model(&rolePermissions).
			Returning("*").
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	var finalEntityRole entities.Role

	err = db.NewSelect().
		Model(&finalEntityRole).
		Relation("Permissions").
		Where("id = ?", entityRole.ID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return entityRole.ToDomain(), nil
}
