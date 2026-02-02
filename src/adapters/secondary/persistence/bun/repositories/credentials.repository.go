package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/core/domain"
)

type CredentialsRepository struct {
	persistenceAdapter *bun.BunAdapter
}

type CredentialsRepositoryDeps struct {
	PersistenceAdapter *bun.BunAdapter
}

func NewCredentialsRepository(deps *CredentialsRepositoryDeps) domain.ICredentialsRepository {
	return &CredentialsRepository{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

func (repository *CredentialsRepository) FindByUsernameOrEmail(
	ctx context.Context,
	usernameOrEmail string,
) (*domain.Credential, error) {
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

	return credential.ToDomain(), nil
}

func (repository *CredentialsRepository) UsernameOrEmailExists(
	ctx context.Context,
	usernameOrEmail string,
) (bool, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var credential entities.Credential
	return db.NewSelect().
		Model(&credential).
		Where("username = ? OR email = ?", usernameOrEmail).
		Exists(ctx)
}

func (repository *CredentialsRepository) CreateOne(ctx context.Context, domainCredential *domain.Credential) (*domain.Credential, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var credential entities.Credential
	credential.FromDomain(domainCredential)

	err := db.
		NewInsert().
		Returning("*").
		Model(&credential).
		Scan(ctx)

	return credential.ToDomain(), err
}
