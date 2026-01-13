package dtos

type UserListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto `json:"pagination"`
	Users               []*UserOutputDto     `json:"users"`
}
