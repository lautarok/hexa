package repositories

import (
	"context"

	bunPersistence "github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/application/domain"
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
		Limit(limit).
		Offset(skip).
		Model(&entityRoleList).
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
