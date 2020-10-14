package dto

import(
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
type UserRequest struct{
	Email string  `validate:"required,email"`
	Password string
	ConfirmPassword string
}

type UserResponse struct{
	Data struct{
		Email string  
		Password string
		ConfirmPassword string
	}
	IsSuccess bool
}