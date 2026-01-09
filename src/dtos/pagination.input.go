package dtos

type PaginationInput struct {
	Page  int `form:"page" validate:"min=1,max=9999"`
	Limit int `form:"limit" validate:"min=1,max=9999"`
}
