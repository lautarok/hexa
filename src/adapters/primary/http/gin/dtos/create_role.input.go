package dtos

import (
	"errors"
	"strings"

	"github.com/lautarok/hexa/src/core/ports"
	rolesCommand "github.com/lautarok/hexa/src/core/usecases/roles/command"
)

type CreateRoleInputDto struct {
	NameEn      string                     `json:"nameEn" validate:"omitempty,min=2,max=35"`
	NameEs      string                     `json:"nameEs" validate:"omitempty,min=2,max=35"`
	NameFr      string                     `json:"nameFr" validate:"omitempty,min=2,max=35"`
	NamePt      string                     `json:"namePt" validate:"omitempty,min=2,max=35"`
	Permissions []*PermissionAliasInputDto `json:"permissions" validate:"required,min=1,max=125"`
}

func (dto *CreateRoleInputDto) Validate(validation ports.ValidationPort) error {
	if strings.TrimSpace(dto.NameEn) == "" &&
		strings.TrimSpace(dto.NameEs) == "" &&
		strings.TrimSpace(dto.NameFr) == "" &&
		strings.TrimSpace(dto.NamePt) == "" {
		return errors.New("no traductions found in request body")
	}

	return validation.Struct(dto)
}

func (dto *CreateRoleInputDto) ToUsecaseInput() *rolesCommand.CreateRoleUsecaseInput {
	permissions := []*rolesCommand.CreateRoleUsecaseInputPermission{}
	for _, permission := range dto.Permissions {
		permissions = append(permissions, permission.ToUsecaseInput())
	}

	return &rolesCommand.CreateRoleUsecaseInput{
		NameEn:      dto.NameEn,
		NameEs:      dto.NameEs,
		NameFr:      dto.NameFr,
		NamePt:      dto.NamePt,
		Permissions: permissions,
	}
}
