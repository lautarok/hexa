package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Role       Role
	Name       string
	Surname    string
	Credential Credential
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type IUsersRepository interface {
	FindMany(ctx context.Context, skip int, limit int) ([]*User, error)
	CreateOne(ctx context.Context, user *User) (*User, error)
	FindOneByID(ctx context.Context, id uuid.UUID) (*User, error)
	DeleteOne(ctx context.Context, id uuid.UUID) error
	UpdateOne(ctx context.Context, user *User) (*User, error)
}
