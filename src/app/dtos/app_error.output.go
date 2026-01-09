package dto

type AppErrorDto struct {
	Code       string `json:"code"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
