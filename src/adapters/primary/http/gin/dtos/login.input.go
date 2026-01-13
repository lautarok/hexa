package dtos

import "github.com/lautarok/hexa/src/application/ports"

type LoginInputDto struct {
	UsernameOrEmail string `validate:"required,username|email" json:"usernameOrEmail"`
	Password        string `validate:"required,securepassword" json:"password"`
}

func (input *LoginInputDto) Validate(validationAdapter ports.ValidationPort) error {
	return validationAdapter.Struct(input)
}
