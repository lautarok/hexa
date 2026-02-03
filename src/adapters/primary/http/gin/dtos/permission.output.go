package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
)

type PermissionOutputDto struct {
	ID        uuid.UUID `json:"id"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPermissionOutputDto(permissionModel *models.Permission) *PermissionOutputDto {
	return &PermissionOutputDto{
		ID:        permissionModel.ID,
		Alias:     permissionModel.Alias,
		CreatedAt: permissionModel.CreatedAt,
	}
}
