package domain

import "github.com/google/uuid"

type Identity struct {
	SubUserID uuid.UUID
	UserID    uuid.UUID
	Email     string
	Username  string
}
