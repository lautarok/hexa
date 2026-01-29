package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/application/domain"
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
	db := repository.dbAdapter.GetDB(ctx)

	var credential entities.Credential
	err := db.NewSelect().
		Model(&credential).
		Where("credential.username = ?", usernameOrEmail).
		WhereOr("credential.email = ?", usernameOrEmail).
		Relation("User").
		Relation("User.Credential").
		Relation("User.Role").
		Relation("User.Role.Permissions").
		Scan(ctx)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, nil
		}
		return nil, err
	}

	return credential.ToDomain(), err
}

func (repository *CredentialsRepository) UsernameOrEmailExists(
	ctx context.Context,
	usernameOrEmail string,
) (bool, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var credential entities.Credential
	return db.NewSelect().
		Model(&credential).
		Where("username = ? OR email = ?", usernameOrEmail).
		Exists(ctx)
}

func (repository *CredentialsRepository) CreateOne(ctx context.Context, domainCredential *domain.Credential) (*domain.Credential, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var credential entities.Credential
	credential.FromDomain(domainCredential)

	err := db.
		NewInsert().
		Returning("*").
		Model(&credential).
		Scan(ctx)

	return credential.ToDomain(), err
}
