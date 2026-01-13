package dtos

import "github.com/lautarok/hexa/src/application/ports"

type PaginationInputDto struct {
	Page  int `form:"page" validate:"min=1,max=9999"`
	Limit int `form:"limit" validate:"min=1,max=9999"`
}

func (input *PaginationInputDto) Validate(validationAdapter ports.ValidationPort) error {
	if input.Page == 0 {
		input.Page = 1
	}

	if input.Limit == 0 {
		input.Limit = 15
	}

	return validationAdapter.Struct(input)
}
