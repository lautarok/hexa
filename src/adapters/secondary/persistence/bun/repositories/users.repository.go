package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	bunPersistence "github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/uptrace/bun"
)

type UsersRepository struct {
	persistenceAdapter *bunPersistence.BunAdapter
}

type UsersRepositoryDeps struct {
	PersistenceAdapter *bunPersistence.BunAdapter
}

func NewUsersRepository(deps *UsersRepositoryDeps) models.IUsersRepository {
	return &UsersRepository{
		persistenceAdapter: deps.PersistenceAdapter,
	}
}

func (repository *UsersRepository) FindMany(ctx context.Context, skip int, limit int) ([]*models.User, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var userList []*entities.User
	err := db.NewSelect().
		Column("id", "name", "surname", "created_at").
		Limit(limit).
		Offset(skip).
		OrderBy("created_at", bun.OrderDesc).
		Model(&userList).
		Relation("Credential", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Column("id", "username", "email")
		}).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	domainUserList := []*models.User{}
	for _, user := range userList {
		domainUserList = append(domainUserList, user.ToDomainModel())
	}

	return domainUserList, nil
}

func (repository *UsersRepository) CreateOne(ctx context.Context, domainUser *models.User) (*models.User, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var user entities.User
	user.FromDomainModel(domainUser)

	err := db.
		NewInsert().
		Returning("*").
		Model(&user).
		Scan(ctx)

	return user.ToDomainModel(), err
}

func (repository *UsersRepository) FindOneByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var user entities.User

	err := db.NewSelect().
		Model(&user).
		Where(`"user"."id" = ?`, id).
		Relation("Credential").
		Relation("Role").
		Relation("Role.Permissions").
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user.ToDomainModel(), nil
}

func (repository *UsersRepository) DeleteOne(ctx context.Context, id uuid.UUID) error {
	db := repository.persistenceAdapter.GetDB(ctx)

	_, err := db.NewDelete().
		Model((*entities.User)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}

func (repository *UsersRepository) UpdateOne(ctx context.Context, user *models.User) (*models.User, error) {
	db := repository.persistenceAdapter.GetDB(ctx)

	var entityUser entities.User
	entityUser.FromDomainModel(user)

	_, err := db.NewUpdate().
		Model(&entityUser).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	err = db.NewSelect().
		Model(&entityUser).
		WherePK().
		Relation("Credential").
		Relation("Role", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Relation("Permissions")
		}).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return entityUser.ToDomainModel(), nil
}
