package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/app/domain"
)

type CredentialsRepository struct {
	dbAdapter *bun.BunAdapter
}

type CredentialsRepositoryDeps struct {
	DBAdapter *bun.BunAdapter
}

func NewCredentialsRepository(deps *CredentialsRepositoryDeps) domain.ICredentialsRepository {
	return &CredentialsRepository{
		dbAdapter: deps.DBAdapter,
	}
}

func (repository *CredentialsRepository) FindByUsernameOrEmail(
	ctx context.Context,
	usernameOrEmail string,
) (*domain.Credential, error) {
	return &domain.Credential{}, nil
}

func (repository *CredentialsRepository) CreateOne(ctx context.Context, domainCredential *domain.Credential) (*domain.Credential, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var credential entities.Credential
	credential.FromDomain(domainCredential)

	_, err := db.
		NewInsert().
		Returning("*").
		Model(&credential).
		Exec(ctx)

	return credential.ToDomain(), err
}
