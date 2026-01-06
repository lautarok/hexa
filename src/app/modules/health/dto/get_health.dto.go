package dto

type GetHealthDto struct {
	StatusCode       int    `json:"statusCode"`
	Message          string `json:"message"`
	ProcessingTimeMs int    `json:"processingTimeMs"`
}
