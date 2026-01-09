package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserOutputDto struct {
	ID         uuid.UUID            `json:"id"`
	Name       string               `json:"name"`
	Surname    string               `json:"surname"`
	Username   string               `json:"username"`
	Credential *CredentialOutputDto `json:"credential"`
	CreatedAt  time.Time            `json:"createdAt"`
}
