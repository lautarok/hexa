package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
)

type UserOutputDto struct {
	ID         uuid.UUID            `json:"id"`
	Name       string               `json:"name"`
	Surname    string               `json:"surname"`
	Credential *CredentialOutputDto `json:"credential,omitempty"`
	CreatedAt  time.Time            `json:"createdAt"`
	Role       *RoleOutputDto       `json:"role"`
}

func NewUserOutputDto(userModel *models.User) *UserOutputDto {
	user := &UserOutputDto{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Surname:   userModel.Surname,
		CreatedAt: userModel.CreatedAt,
	}

	user.Credential = NewCredentialOutputDto(&userModel.Credential)
	user.Role = NewRoleOutputDto(&userModel.Role)

	return user
}
