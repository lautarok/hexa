package domain

type AppError struct {
	Code    string
	Message string
}

func (domain *AppError) Error() string {
	return domain.Message
}
