package domain

import "time"

type Credential struct {
	ID           string
	Email        string
	Username     string
	PasswordHash string
	User         *User
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
