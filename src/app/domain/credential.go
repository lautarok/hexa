package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID           uuid.UUID `validation:"uuid"`
	Email        string    `validation:"email"`
	Username     string    `validation:"username"`
	PasswordHash string    `validation:"required"`
	User         *User     `validation:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ICredentialsRepository interface {
	FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*Credential, error)
	CreateOne(ctx context.Context, credentials *Credential) (*Credential, error)
}
