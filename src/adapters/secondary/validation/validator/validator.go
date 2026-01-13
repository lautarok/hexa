package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/lautarok/hexa/src/application/ports"
)

type ValidatorAdapter struct {
	validator *validator.Validate
}

func NewValidatorAdapter() ports.ValidationPort {
	v := validator.New()

	v.RegisterValidation("securepassword", func(field validator.FieldLevel) bool {
		password := field.Field().String()

		var (
			hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString(password)
			hasLower   = regexp.MustCompile(`[a-z]`).MatchString(password)
			hasNumber  = regexp.MustCompile(`[0-9]`).MatchString(password)
			hasSpecial = regexp.MustCompile(`[@$!%*?&]`).MatchString(password)
		)

		return hasUpper &&
			hasLower &&
			hasNumber &&
			hasSpecial &&
			len(password) >= 6 &&
			len(password) <= 20
	})

	v.RegisterValidation("username", func(field validator.FieldLevel) bool {
		username := field.Field().String()
		regExp := regexp.MustCompile(
			`^[a-zA-Z0-9._]*$`,
		)
		return regExp.MatchString(username) &&
			len(username) >= 5 &&
			len(username) <= 25
	})

	return &ValidatorAdapter{
		validator: v,
	}
}

func (adapter *ValidatorAdapter) Struct(structure any) error {
	return adapter.validator.Struct(structure)
}
