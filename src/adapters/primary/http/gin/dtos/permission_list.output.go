package dtos

type PermissionListOutputDto struct {
	PaginationOutputDto *PaginationOutputDto   `json:"pagination"`
	Permissions         []*PermissionOutputDto `json:"permissions"`
}
