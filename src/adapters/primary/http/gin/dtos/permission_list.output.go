package dtos

import "github.com/lautarok/hexa/src/domain/models"

type PermissionListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto   `json:"pagination"`
	Permissions         []*PermissionOutputDto `json:"permissions"`
}

func NewPermissionListOutputDto(pagination *PaginationInputDto, permissionModels []*models.Permission) *PermissionListOutputDto {
	permissions := []*PermissionOutputDto{}
	for _, permission := range permissionModels {
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
