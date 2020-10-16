package dto

import(
	"github.com/go-playground/validator/v10"

)

var validate *validator.Validate
type RegistrationRequest struct{
	Email string  `validate:"required,email"`
	Password string `validate:"passwd"`
	ConfirmPassword string
}

type RegistrationResponse struct{
	Data struct{
		Id uint
		Email string  
	}
	IsSuccess bool
}


