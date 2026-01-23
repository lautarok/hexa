package dtos

import "github.com/lautarok/hexa/src/application/ports"

type SignupInputDto struct {
	Name           string `validate:"required,min=3,max=40" json:"name"`
	Surname        string `validate:"required,min=3,max=40" json:"surname"`
	Username       string `validate:"required,username" json:"username"`
	Email          string `validate:"required,email" json:"email"`
	Password       string `validate:"required,securepassword" json:"password"`
	RepeatPassword string `validate:"required,securepassword,eqfield=Password" json:"repeatPassword"`
}

func (dto *SignupInputDto) Validate(validationAdapter ports.ValidationPort) error {
	return validationAdapter.Struct(dto)
}
