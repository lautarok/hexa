package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/application/domain"
	bunInfra "github.com/uptrace/bun"
)

type PermissionsRepository struct {
	dbAdapter *bun.BunAdapter
}

type PermissionsRepositoryDeps struct {
	DBAdapter *bun.BunAdapter
}

func NewPermissionsRepository(deps *PermissionsRepositoryDeps) domain.IPermissionsRepository {
	return &PermissionsRepository{
		dbAdapter: deps.DBAdapter,
	}
}

func (repository *PermissionsRepository) FindMany(ctx context.Context, skip int, limit int) ([]*domain.Permission, error) {
	db := repository.dbAdapter.GetDB(ctx)

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

	domainPermissions := []*domain.Permission{}
	for _, permission := range entityPermissions {
		domainPermissions = append(domainPermissions, permission.ToDomain())
	}

	return domainPermissions, nil
}

func (repository *PermissionsRepository) FindManyByAlias(ctx context.Context, aliases ...string) ([]*domain.Permission, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var entityPermissions []*entities.Permission
	err := db.NewSelect().
		Model(&entityPermissions).
		Where("alias IN (?)", bunInfra.In(aliases)).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var domainPermissions []*domain.Permission
	for _, permission := range entityPermissions {
		domainPermissions = append(domainPermissions, permission.ToDomain())
	}

	return domainPermissions, nil
}
