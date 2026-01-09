package dtos

import "github.com/google/uuid"

type CredentialOutputDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
