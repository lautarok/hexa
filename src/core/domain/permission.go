package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID
	Alias     string
	Roles     []*Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IPermissionsRepository interface {
	FindMany(ctx context.Context, skip int, limit int) ([]*Permission, error)
	FindManyByAlias(ctx context.Context, aliases ...string) ([]*Permission, error)
}
