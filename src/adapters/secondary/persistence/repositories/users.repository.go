package repositories

import (
	"context"

	"github.com/lautarok/hexa/src/adapters/secondary/persistence/entities"
	"github.com/lautarok/hexa/src/app/domain"
	"github.com/uptrace/bun"
)

type UsersRepository struct {
	db bun.IDB
}

type UsersRepositoryDeps struct {
	DB bun.IDB
}

func NewUsersRepository(deps *UsersRepositoryDeps) domain.IUsersRepository {
	return &UsersRepository{
		db: deps.DB,
	}
}

func (repository *UsersRepository) FindMany(ctx context.Context, take int, limit int) ([]*domain.User, error) {
	var userList []*entities.User
	err := repository.db.NewSelect().
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
