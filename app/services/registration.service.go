package services

import(
	"go-connect/app/dto"
	"go-connect/app/models"
	"errors"
)

func Registration(userReq dto.RegistrationRequest)  (dto.RegistrationResponse, error){

	var userModel = models.User{}
	var registrationResponse = dto.RegistrationResponse{}
	var err error = nil
	var isConfirmPasswordMatched = confirmPasswordValidator(userReq.Password,userReq.ConfirmPassword)
	isError := false
	if isConfirmPasswordMatched {
		userModel.Email = userReq.Email
		userModel.Password  = userReq.Password
		result := db.Create(&userModel)
		if result != nil {
			registrationResponse.Data.Email = userModel.Email
			registrationResponse.Data.Id = userModel.ID
			registrationResponse.IsSuccess = true
		}else{
			isError = true
		}	
	}else{
		isError = true
	}

	if isError {
		err = errors.New("user registration failed")
	}
	return registrationResponse,err
}


func confirmPasswordValidator(password string, confirmPassword string) bool{
		return password == confirmPassword
}