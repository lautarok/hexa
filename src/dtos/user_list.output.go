package dtos

type UserListOutputDto struct {
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
	Users []*UserOutputDto `json:"users"`
}
