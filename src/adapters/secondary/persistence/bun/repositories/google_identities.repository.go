package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/domain/models"
)

type GoogleIdentitiesRepository struct {
	persistenceAdapter *bun.BunAdapter
}

type GoogleIdentitiesRepositoryDeps struct {
	PersistenceAdapter *bun.BunAdapter
}

func NewGoogleIdentitiesRepository(deps *GoogleIdentitiesRepositoryDeps) models.IGoogleIdentitiesRepository {
	return &GoogleIdentitiesRepository{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

func (repository *GoogleIdentitiesRepository) FindOneByGoogleID(ctx context.Context, googleId string) (*models.GoogleIdentity, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var googleIdentity models.GoogleIdentity

	err := db.
		NewSelect().
		Model(&googleIdentity).
		Where("google_id = ?", googleId).
		Scan(ctx)

	return &googleIdentity, err
}

func (repository *GoogleIdentitiesRepository) CreateOne(ctx context.Context, googleIdentity *models.GoogleIdentity) (*models.GoogleIdentity, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	entity := entities.GoogleIdentity{}
	entity.FromDomainModel(googleIdentity)

	err := db.
		NewInsert().
		Model(&entity).
		Returning("*").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return entity.ToDomainModel(), nil
}
