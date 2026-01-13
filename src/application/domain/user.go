package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID   `validation:"required,uuid"`
	Role       *Role       `validation:"required"`
	Name       string      `validation:"required,min=3,max=40"`
	Surname    string      `validation:"required,min=3,max=40"`
	Credential *Credential `validation:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type IUsersRepository interface {
	FindMany(ctx context.Context, skip int, limit int) ([]*User, error)
	CreateOne(ctx context.Context, domainUser *User) (*User, error)
	FindOneByID(ctx context.Context, id uuid.UUID) (*User, error)
}
