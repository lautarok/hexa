package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/domain/models"
	bunInfra "github.com/uptrace/bun"
)

type PermissionsRepository struct {
	persistenceAdapter *bun.BunAdapter
}

type PermissionsRepositoryDeps struct {
	PersistenceAdapter *bun.BunAdapter
}

func NewPermissionsRepository(deps *PermissionsRepositoryDeps) models.IPermissionsRepository {
	return &PermissionsRepository{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

func (repository *PermissionsRepository) FindMany(ctx context.Context, skip int, limit int) ([]*models.Permission, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var entityPermissions []*entities.Permission
	err := db.NewSelect().
		Limit(limit).
		Offset(skip).
		OrderBy("created_at", bunInfra.OrderAsc).
		Model(&entityPermissions).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	domainPermissions := []*models.Permission{}
	for _, permission := range entityPermissions {
		domainPermissions = append(domainPermissions, permission.ToDomainModel())
	}

	return domainPermissions, nil
}

func (repository *PermissionsRepository) FindManyByAlias(ctx context.Context, aliases ...string) ([]*models.Permission, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var entityPermissions []*entities.Permission
	err := db.NewSelect().
		Model(&entityPermissions).
		Where("alias IN (?)", bunInfra.In(aliases)).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var domainPermissions []*models.Permission
	for _, permission := range entityPermissions {
		domainPermissions = append(domainPermissions, permission.ToDomainModel())
	}

	return domainPermissions, nil
}
