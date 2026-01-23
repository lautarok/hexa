package dtos

import (
	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/application/domain"
)

type CredentialOutputDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func NewCredentialOutputDto(domainCredential *domain.Credential) *CredentialOutputDto {
	credential := &CredentialOutputDto{}

	credential.ID = domainCredential.ID
	credential.Username = domainCredential.Username
	credential.Email = domainCredential.Email

	return credential
}
