package domain

import "time"

type User struct {
	ID         string      `validate:"required,uuid"`
	Role       *Role       `validate:"required"`
	Name       string      `validate:"required,min=3,max=40"`
	Surname    string      `validate:"required,min=3,max=40"`
	Credential *Credential `validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
