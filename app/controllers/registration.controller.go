package controllers

import (
	"net/http"
	"go-connect/app/utility"
	"github.com/gorilla/context"
	"go-connect/app/dto"
	"go-connect/app/services"	
)


var Register = func(w http.ResponseWriter, r *http.Request){ 
	logger.Info.Println("test...")
	body := context.Get(r,"reqBody")
	registrationReq, ok := body.(dto.RegistrationRequest)
	registrationResponse := dto.RegistrationResponse{} 
	var err error = nil
    logger.Info.Println("ok.........",ok)
	if ok {
		logger.Info.Println("registrationReq.........",registrationReq)
		registrationResponse, err = services.Registration(registrationReq)
		logger.Info.Println(err)
	}
	var data = make(map[string]interface{})
	data["Data"] = registrationResponse
	utility.Respond(w, data)
		
}