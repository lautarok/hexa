package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/domain/models"
)

type CredentialsRepository struct {
	persistenceAdapter *bun.BunAdapter
}

type CredentialsRepositoryDeps struct {
	PersistenceAdapter *bun.BunAdapter
}

func NewCredentialsRepository(deps *CredentialsRepositoryDeps) models.ICredentialsRepository {
	return &CredentialsRepository{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

func (repository *CredentialsRepository) FindByUsernameOrEmail(
	ctx context.Context,
	usernameOrEmail string,
) (*models.Credential, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var credential entities.Credential
	err := db.NewSelect().
		Model(&credential).
		Where("credential.username = ?", usernameOrEmail).
		WhereOr("credential.email = ?", usernameOrEmail).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return credential.ToDomainModel(), nil
}

func (repository *CredentialsRepository) CreateOne(ctx context.Context, credentialModel *models.Credential) (*models.Credential, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var credential entities.Credential
	credential.FromDomainModel(credentialModel)

	err := db.
		NewInsert().
		Returning("*").
		Model(&credential).
		Scan(ctx)

	return credential.ToDomainModel(), err
}

func (repository *CredentialsRepository) UpdateOne(ctx context.Context, credentialModel *models.Credential) (*models.Credential, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var credential entities.Credential
	credential.FromDomainModel(credentialModel)

	err := db.
		NewUpdate().
		Where("user_id = ?", credentialModel.UserID).
		Model(&credential).
		Returning("*").
		Scan(ctx)

	return credential.ToDomainModel(), err
}
