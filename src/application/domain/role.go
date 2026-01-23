package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID     `validation:"uuid"`
	NameFr      string        `validation:"min=2,max=30"`
	NameEs      string        `validation:"min=2,max=30"`
	NameEn      string        `validation:"min=2,max=30"`
	NamePt      string        `validation:"min=2,max=30"`
	Permissions []*Permission `validation:"required,min=1"`
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
