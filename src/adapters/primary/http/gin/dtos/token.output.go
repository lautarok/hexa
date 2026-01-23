package dtos

import "github.com/lautarok/hexa/src/application/domain"

type TokenOutputDto struct {
	Token string        `json:"token"`
	Exp   int64         `json:"expirationTime"`
	User  UserOutputDto `json:"user"`
}

func NewTokenOutputDto(token string, exp int64, user *domain.User) *TokenOutputDto {
	return &TokenOutputDto{
		Token: token,
		Exp:   exp,
		User:  *NewUserOutputDto(user),
	}
}
