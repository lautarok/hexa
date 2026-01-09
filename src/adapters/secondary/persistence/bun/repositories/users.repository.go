package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun"
	"github.com/lautarok/hexa/src/adapters/secondary/persistence/bun/entities"
	"github.com/lautarok/hexa/src/app/domain"
)

type UsersRepository struct {
	dbAdapter *bun.BunAdapter
}

type UsersRepositoryDeps struct {
	DBAdapter *bun.BunAdapter
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
		Model(&userList).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var domainUserList []*domain.User
	for _, user := range userList {
		domainUserList = append(domainUserList, user.ToDomain())
	}

	if domainUserList == nil {
		return []*domain.User{}, nil
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
