package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	bunPersistence "github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/application/domain"
	"github.com/uptrace/bun"
)

type UsersRepository struct {
	dbAdapter *bunPersistence.BunAdapter
}

type UsersRepositoryDeps struct {
	DBAdapter *bunPersistence.BunAdapter
}

func NewUsersRepository(deps *UsersRepositoryDeps) domain.IUsersRepository {
	return &UsersRepository{
		dbAdapter: deps.DBAdapter,
	}
}

func (repository *UsersRepository) FindMany(ctx context.Context, skip int, limit int) ([]*domain.User, error) {
	db := repository.dbAdapter.GetDB(ctx)

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

	domainUserList := []*domain.User{}
	for _, user := range userList {
		domainUserList = append(domainUserList, user.ToDomain())
	}

	return domainUserList, nil
}

func (repository *UsersRepository) CreateOne(ctx context.Context, domainUser *domain.User) (*domain.User, error) {
	db := repository.dbAdapter.GetDB(ctx)

	var user entities.User
	user.FromDomain(domainUser)

	_, err := db.
		NewInsert().
		Returning("*").
		Model(&user).
		Exec(ctx)

	return user.ToDomain(), err
}

func (respository *UsersRepository) FindOneByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	db := respository.dbAdapter.GetDB(ctx)

	var user entities.User

	err := db.NewSelect().
		Model(&user).
		Where(`"user"."id" = ?`, id).
		Relation("Credential").
		Relation("Role").
		Relation("Role.Permissions").
		Scan(ctx)

	log.Println(user.Role)

	return user.ToDomain(), err
}
