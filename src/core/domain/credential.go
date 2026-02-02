package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ICredentialsRepository interface {
	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*Credential, error)
	UsernameOrEmailExists(ctx context.Context, usernameOrEmail string) (bool, error)
	CreateOne(ctx context.Context, credentials *Credential) (*Credential, error)
}
