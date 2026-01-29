package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/application/domain"
)

type RoleOutputDto struct {
	ID          uuid.UUID              `json:"id"`
	NameEn      string                 `json:"nameEn"`
	NameEs      string                 `json:"nameEs"`
	NameFr      string                 `json:"nameFr"`
	NamePt      string                 `json:"namePt"`
	NameNl      string                 `json:"nameNl"`
	Permissions []*PermissionOutputDto `json:"permissions"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

func NewRoleOutputDto(domainRole *domain.Role) *RoleOutputDto {
	role := &RoleOutputDto{
		ID:        domainRole.ID,
		NameEn:    domainRole.NameEn,
		NameEs:    domainRole.NameEs,
		NameFr:    domainRole.NameFr,
		NamePt:    domainRole.NamePt,
		NameNl:    domainRole.NameNl,
		CreatedAt: domainRole.CreatedAt,
		UpdatedAt: domainRole.UpdatedAt,
	}

	if domainRole.Permissions != nil {
		for _, permission := range domainRole.Permissions {
			role.Permissions = append(role.Permissions, NewPermissionOutputDto(permission))
		}
	}

	return role
}
