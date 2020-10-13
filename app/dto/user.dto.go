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