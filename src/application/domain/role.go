package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID
	NameFr      string
	NameEs      string
	NameEn      string
	NamePt      string
	NameNl      string
	Permissions []*Permission
	Users       []*User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type IRolesRepository interface {
	GetOneByAlias(ctx context.Context, alias string) (*Role, error)
	FindMany(ctx context.Context, skip int, limit int) ([]*Role, error)
	CreateOne(ctx context.Context, role *Role) (*Role, error)
	DeleteOne(ctx context.Context, id uuid.UUID) error
	UpdateOne(ctx context.Context, role *Role) (*Role, error)
}
