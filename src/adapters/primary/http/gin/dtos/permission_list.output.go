package dtos

import "github.com/lautarok/hexa/src/core/domain"

type PermissionListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto   `json:"pagination"`
	Permissions         []*PermissionOutputDto `json:"permissions"`
}

func NewPermissionListOutputDto(pagination *PaginationInputDto, domainPermissions []*domain.Permission) *PermissionListOutputDto {
	permissions := []*PermissionOutputDto{}
	for _, permission := range domainPermissions {
		permissions = append(permissions, NewPermissionOutputDto(permission))
	}

	return &PermissionListOutputDto{
		PaginationOutputDto: &PaginationOutputDto{
			Page:  pagination.Page,
			Limit: pagination.Limit,
		},
		Permissions: permissions,
	}
}
