package dtos

import "github.com/lautarok/hexa/src/domain/models"

type UserListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto `json:"pagination"`
	Users               []*UserOutputDto     `json:"users"`
}

func NewUserListOutputDto(pagination *PaginationInputDto, userModels []*models.User) *UserListOutputDto {
	users := []*UserOutputDto{}

	for _, user := range userModels {
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
