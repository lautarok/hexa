package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/core/domain"
)

type UserOutputDto struct {
	ID         uuid.UUID            `json:"id"`
	Name       string               `json:"name"`
	Surname    string               `json:"surname"`
	Credential *CredentialOutputDto `json:"credential,omitempty"`
	CreatedAt  time.Time            `json:"createdAt"`
	Role       *RoleOutputDto       `json:"role"`
}

func NewUserOutputDto(domainUser *domain.User) *UserOutputDto {
	user := &UserOutputDto{
		ID:        domainUser.ID,
		Name:      domainUser.Name,
		Surname:   domainUser.Surname,
		CreatedAt: domainUser.CreatedAt,
	}

	user.Credential = NewCredentialOutputDto(&domainUser.Credential)
	user.Role = NewRoleOutputDto(&domainUser.Role)

	return user
}
