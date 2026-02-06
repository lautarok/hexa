package dtos

import (
	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/ports"
)

type UpdateUserCredentialInputDto struct {
	UserID   uuid.UUID `json:"userId" validate:"required,uuid"`
	Username string    `json:"username" validate:"username"`
}

func (dto *UpdateUserCredentialInputDto) Validate(validationAdapter ports.ValidationPort) error {
	return validationAdapter.Struct(dto)
}
