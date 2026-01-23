package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/application/domain"
)

type PermissionOutputDto struct {
	ID        uuid.UUID `json:"id"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPermissionOutputDto(domainPermission *domain.Permission) *PermissionOutputDto {
	return &PermissionOutputDto{
		ID:        domainPermission.ID,
		Alias:     domainPermission.Alias,
		CreatedAt: domainPermission.CreatedAt,
	}
}
