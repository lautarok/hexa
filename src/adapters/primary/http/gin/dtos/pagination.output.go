package dtos

type PaginationOutputDto struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewPaginationOutputDto(pagination *PaginationInputDto) *PaginationOutputDto {
	return &PaginationOutputDto{
		Page:  pagination.Page,
		Limit: pagination.Limit,
	}
}
