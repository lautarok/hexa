package models

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
	CreateOne(ctx context.Context, credentials *Credential) (*Credential, error)
}
