package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	bunPersistence "github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/uptrace/bun"
)

type RolesRepository struct {
	persistenceAdapter *bunPersistence.BunAdapter
}

type RolesRepositoryDeps struct {
	PersistenceAdapter *bunPersistence.BunAdapter
}

func NewRolesRepository(deps *RolesRepositoryDeps) models.IRolesRepository {
	return &RolesRepository{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

func (repository *RolesRepository) GetOneBySlug(ctx context.Context, slug string) (*models.Role, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var role entities.Role
	err := db.
		NewSelect().
		Model(&role).
		Where("slug = ?", slug).
		Scan(ctx)

	return role.ToDomainModel(), err
}

func (repository *RolesRepository) FindMany(ctx context.Context, skip int, limit int) ([]*models.Role, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

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

	roleList := []*models.Role{}

	for _, entityRole := range entityRoleList {
		roleList = append(roleList, entityRole.ToDomainModel())
	}

	return roleList, err
}

func (repository *RolesRepository) CreateOne(ctx context.Context, domainRole *models.Role) (*models.Role, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var entityRole entities.Role
	entityRole.FromDomainModel(domainRole)

	err := db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		err := tx.NewInsert().
			Model(&entityRole).
			Returning("*").
			Scan(ctx)
		if err != nil {
			return err
		}

		rolePermissions := []*entities.RolePermission{}

		for _, permission := range domainRole.Permissions {
			rolePermissions = append(rolePermissions, &entities.RolePermission{
				RoleID:       entityRole.ID,
				PermissionID: permission.ID,
			})
		}

		err = tx.NewInsert().
			Model(&rolePermissions).
			Returning("*").
			Scan(ctx)
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

	return entityRole.ToDomainModel(), nil
}

func (repository *RolesRepository) DeleteOne(ctx context.Context, id uuid.UUID) error {
	db := repository.persistenceAdapter.GetDB(ctx)

	_, err := db.NewDelete().
		Model((*entities.Role)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}

func (repository *RolesRepository) UpdateOne(ctx context.Context, role *models.Role) (*models.Role, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var entityRole *entities.Role
	entityRole.FromDomainModel(role)

	err := db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewUpdate().
			Model(&entityRole).
			WherePK().
			Exec(ctx)
		if err != nil {
			return err
		}

		if role.Permissions == nil {
			return nil
		}

		_, err = tx.NewDelete().
			Model((*entities.RolePermission)(nil)).
			Where("role_id = ?", entityRole.ID).
			Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewInsert().
			Model(&entityRole.Permissions).
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = db.NewSelect().
		Model(&entityRole).
		WherePK().
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return entityRole.ToDomainModel(), nil
}
