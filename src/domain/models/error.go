package models

type AppError struct {
	Code    string
	Message string
}

func (model *AppError) Error() string {
	return model.Message
}
