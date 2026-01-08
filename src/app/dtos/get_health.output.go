package dto

type GetHealthOutputDto struct {
	StatusCode       int    `json:"statusCode"`
	Message          string `json:"message"`
	ProcessingTimeMs int    `json:"processingTimeMs"`
}
