package dtos

import (
	rolesCommand "github.com/lautarok/hexa/src/application/usecases/roles/command"
	"github.com/lautarok/hexa/src/domain/ports"
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
