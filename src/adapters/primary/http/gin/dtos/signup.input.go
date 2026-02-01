package dtos

import (
	"strings"

	"github.com/lautarok/hexa/src/application/ports"
)

type SignupInputDto struct {
	Name           string `validate:"required,min=3,max=40,name" json:"name"`
	Surname        string `validate:"required,min=3,max=40,name" json:"surname"`
	Username       string `validate:"required,username" json:"username"`
	Email          string `validate:"required,email" json:"email"`
	Password       string `validate:"required,securepassword" json:"password"`
	RepeatPassword string `validate:"required,securepassword,eqfield=Password" json:"repeatPassword"`
}

func (dto *SignupInputDto) normalizeName(input string) string {
	input = strings.TrimSpace(input)
	input = strings.Join(strings.Fields(input), " ")

	words := strings.Split(input, " ")
	lowerParticles := map[string]bool{
		"de":  true,
		"da":  true,
		"do":  true,
		"van": true,
		"von": true,
		"der": true,
		"la":  true,
		"le":  true,
		"del": true,
	}

	for i, word := range words {
		lowerWord := strings.ToLower(word)

		if i > 0 && lowerParticles[lowerWord] {
			words[i] = lowerWord
			continue
		}

		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " ")
}

func (dto *SignupInputDto) Validate(validationAdapter ports.ValidationPort) error {
	dto.Name = dto.normalizeName(dto.Name)
	dto.Surname = dto.normalizeName(dto.Surname)
	return validationAdapter.Struct(dto)
}
