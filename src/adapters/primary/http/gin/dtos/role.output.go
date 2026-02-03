package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/models"
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

func NewRoleOutputDto(roleModel *models.Role) *RoleOutputDto {
	role := &RoleOutputDto{
		ID:        roleModel.ID,
		NameEn:    roleModel.NameEn,
		NameEs:    roleModel.NameEs,
		NameFr:    roleModel.NameFr,
		NamePt:    roleModel.NamePt,
		NameNl:    roleModel.NameNl,
		CreatedAt: roleModel.CreatedAt,
		UpdatedAt: roleModel.UpdatedAt,
	}

	if roleModel.Permissions != nil {
		for _, permission := range roleModel.Permissions {
			role.Permissions = append(role.Permissions, NewPermissionOutputDto(&permission))
		}
	}

	return role
}
