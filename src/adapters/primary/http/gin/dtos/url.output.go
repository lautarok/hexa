package dtos

type URLOutputDto struct {
	URL string `json:"url"`
}

func NewURLOutputDto(url string) *URLOutputDto {
	return &URLOutputDto{
		URL: url,
	}
}
