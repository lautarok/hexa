package dtos

type RoleListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto `json:"pagination"`
	Roles               []*RoleOutputDto     `json:"roles"`
}
