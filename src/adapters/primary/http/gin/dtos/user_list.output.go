package dtos

import "github.com/lautarok/hexa/src/core/domain"

type UserListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto `json:"pagination"`
	Users               []*UserOutputDto     `json:"users"`
}

func NewUserListOutputDto(pagination *PaginationInputDto, domainUsers []*domain.User) *UserListOutputDto {
	users := []*UserOutputDto{}

	for _, user := range domainUsers {
		users = append(users, NewUserOutputDto(user))
	}

	return &UserListOutputDto{
		PaginationOutputDto: &PaginationOutputDto{
			Page:  pagination.Page,
			Limit: pagination.Limit,
		},
		Users: users,
	}
}
