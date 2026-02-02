package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/core/domain"
)

type GoogleIdentitiesRepository struct {
	persistenceAdapter *bun.BunAdapter
}

type GoogleIdentitiesRepositoryDeps struct {
	PersistenceAdapter *bun.BunAdapter
}

func NewGoogleIdentitiesRepository(deps *GoogleIdentitiesRepositoryDeps) domain.IGoogleIdentitiesRepository {
	return &GoogleIdentitiesRepository{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

func (repository *GoogleIdentitiesRepository) FindOneByGoogleID(ctx context.Context, googleId string) (*domain.GoogleIdentity, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var googleIdentity *domain.GoogleIdentity

	err := db.
		NewSelect().
		Model(&googleIdentity).
		Where("google_id = ?", googleId).
		Relation("User").
		Relation("User.Credential").
		Relation("User.Role").
		Relation("User.Role.Permissions").
		Scan(ctx)

	return googleIdentity, err
}

func (repository *GoogleIdentitiesRepository) CreateOne(ctx context.Context, googleIdentity *domain.GoogleIdentity) error {
	db := repository.persistenceAdapter.GetDB(ctx)

	var entity *entities.GoogleIdentity
	entity.FromDomain(googleIdentity)

	_, err := db.
		NewInsert().
		Model(&entity).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
