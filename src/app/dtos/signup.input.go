package dtos

type SignupInputDto struct {
	Name           string `validate:"required,min=3,max=40"`
	Surname        string `validate:"required,min=3,max=40"`
	Username       string `validate:"required,min=5,max=25,username"`
	Email          string `validate:"required,email"`
	Password       string `validate:"required,min=6,max=20,securepassword"`
	RepeatPassword string `validate:"required,min=6,max=20,securepassword,eqfield=Password"`
}
