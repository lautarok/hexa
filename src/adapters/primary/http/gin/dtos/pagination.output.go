package dtos

type PaginationOutputDto struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
