package dtos

import "github.com/lautarok/hexa/src/application/domain"

type RoleListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto `json:"pagination"`
	Roles               []*RoleOutputDto     `json:"roles"`
}

func NewRoleListOutputDto(pagination *PaginationInputDto, domainRoles []*domain.Role) *RoleListOutputDto {
	roles := []*RoleOutputDto{}

	for _, role := range domainRoles {
		roles = append(roles, NewRoleOutputDto(role))
	}

	return &RoleListOutputDto{
		PaginationOutputDto: NewPaginationOutputDto(pagination),
		Roles:               roles,
	}
}
