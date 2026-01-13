package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserOutputDto struct {
	ID         uuid.UUID            `json:"id"`
	Name       string               `json:"name"`
	Surname    string               `json:"surname"`
	Credential *CredentialOutputDto `json:"credential,omitempty"`
	CreatedAt  time.Time            `json:"createdAt"`
}
