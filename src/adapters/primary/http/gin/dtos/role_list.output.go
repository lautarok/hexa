package dtos

import "github.com/lautarok/hexa/src/domain/models"

type RoleListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto `json:"pagination"`
	Roles               []*RoleOutputDto     `json:"roles"`
}

func NewRoleListOutputDto(pagination *PaginationInputDto, roleModels []*models.Role) *RoleListOutputDto {
	roles := []*RoleOutputDto{}

	for _, role := range roleModels {
		roles = append(roles, NewRoleOutputDto(role))
	}

	return &RoleListOutputDto{
		PaginationOutputDto: NewPaginationOutputDto(pagination),
		Roles:               roles,
	}
}
