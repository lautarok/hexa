package dtos

import (
	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
)

type CredentialOutputDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func NewCredentialOutputDto(credentialModel *models.Credential) *CredentialOutputDto {
	credential := &CredentialOutputDto{}

	credential.ID = credentialModel.ID
	credential.Username = credentialModel.Username
	credential.Email = credentialModel.Email

	return credential
}
