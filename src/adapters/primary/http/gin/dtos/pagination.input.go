package dtos

import "github.com/lautarok/hexa/src/domain/ports"

type PaginationInputDto struct {
	Page  int `form:"page" validate:"min=1,max=9999"`
	Limit int `form:"limit" validate:"min=1,max=9999"`
}

func (dto *PaginationInputDto) Validate(validationAdapter ports.ValidationPort) error {
	if dto.Page == 0 {
		dto.Page = 1
	}

	if dto.Limit == 0 {
		dto.Limit = 15
	}

	return validationAdapter.Struct(dto)
}
