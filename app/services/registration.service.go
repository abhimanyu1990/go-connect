package services

import(
	"go-connect/app/dto"
	"go-connect/app/models"
	"go-connect/app/utility"
	"errors"
	"crypto/rand"
	"encoding/hex"
)

// function to handle user registration process
// verify the password and confirm password is same or not
// save user into the database
// call func generateRegistrationTokenAndSendEmail to generate registration token and verification link on email
func Registration(userReq dto.RegistrationRequest)  (dto.RegistrationResponse, error){

	var userModel = models.User{}
	var registrationResponse = dto.RegistrationResponse{}
	var err error = nil
	var errorMessage string = ""
	salt,err := generateSalt()
	if userReq.Password == userReq.ConfirmPassword {
		userModel.Email = userReq.Email
		password := []byte(userReq.Password)
		userModel.Password, err  = utility.Encrypt([] byte(salt), password)
		userModel.Salt = salt
		result := db.Create(&userModel)
		if result != nil {
			registrationResponse.Data.Email = userModel.Email
			registrationResponse.Data.Id = userModel.Id
			registrationResponse.IsSuccess = true
			go generateRegistrationTokenAndSendEmail(userModel.Email)
		}else{
			errorMessage = "Database Error"
		}	
	}else{
		errorMessage = "ConfirmPassword and Password field do not match"
	}

	if err != nil || errorMessage != "" {
		err = errors.New(errorMessage)
	}
	return registrationResponse,err
}

// use to generate salt which is used to encrypt user password
func generateSalt() (string, error){
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	salt := hex.EncodeToString(bytes)
	return salt,err
}

// use to generate token
func generateToken() (string, error){
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	token := hex.EncodeToString(bytes)
	return token,err
}

// Send  registration verifcation link to registered user email
func generateRegistrationTokenAndSendEmail(toEmail string) bool{
	host := env["app_url"]
	port := env["app_port"] 
	isMailSent := true
	token,err := generateToken()
	logger.Error.Println(err)
	result := db.Create(&models.RegistrationToken{Email:toEmail,Token:token})
	if result != nil {
		message := " Please click on this link for verifiction "+host+":"+port+"/"+token
		isMailSent = SendEmail(message,toEmail,"Registration  Verification")
	}
	return isMailSent
}

