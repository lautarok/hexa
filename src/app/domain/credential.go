package domain

import "time"

type Credential struct {
	ID           string `validate:"uuid"`
	Email        string `validate:"email"`
	Username     string `validate:"username"`
	PasswordHash string `validate:"required"`
	User         *User  `validate:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
