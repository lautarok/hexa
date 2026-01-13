package dtos

type TokenOutputDto struct {
	Token string `json:"token"`
	Exp   int64  `json:"expirationTime"`
}
