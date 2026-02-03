package dtos

import "github.com/lautarok/hexa/src/domain/ports"

type LoginInputDto struct {
	UsernameOrEmail string `validate:"required,username|email" json:"usernameOrEmail"`
	Password        string `validate:"required,securepassword" json:"password"`
}

func (dto *LoginInputDto) Validate(validationAdapter ports.ValidationPort) error {
	return validationAdapter.Struct(dto)
}
