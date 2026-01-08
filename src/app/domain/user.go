package domain

import (
	"context"
	"time"
)

type User struct {
	ID         string      `validate:"required,uuid"`
	Role       *Role       `validate:"required"`
	Name       string      `validate:"required,min=3,max=40"`
	Surname    string      `validate:"required,min=3,max=40"`
	Credential *Credential `validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type IUsersRepository interface {
	FindMany(ctx context.Context, take int, limit int) ([]*User, error)
}
