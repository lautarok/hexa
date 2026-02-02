package dtos

import (
	"github.com/lautarok/hexa/src/core/ports"
	rolesCommand "github.com/lautarok/hexa/src/core/usecases/roles/command"
)

type PermissionAliasInputDto struct {
	Alias string `json:"alias" validate:"required,min=2,max=150"`
}

func (dto *PermissionAliasInputDto) Validate(validation ports.ValidationPort) error {
	return validation.Struct(dto)
}

func (dto *PermissionAliasInputDto) ToUsecaseInput() *rolesCommand.CreateRoleUsecaseInputPermission {
	return &rolesCommand.CreateRoleUsecaseInputPermission{
		Alias: dto.Alias,
	}
}
