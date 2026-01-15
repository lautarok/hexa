package dtos

import "github.com/lautarok/hexa/src/application/ports"

type PermissionAliasInputDto struct {
	Alias string `json:"alias" validate:"required,min=2,max=150"`
}

func (dto *PermissionAliasInputDto) Validate(validation ports.ValidationPort) error {
	return validation.Struct(dto)
}
