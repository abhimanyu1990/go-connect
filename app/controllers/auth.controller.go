package controllers

import (
	"net/http"
	"go-connect/app/utility"
)

var Register = func(w http.ResponseWriter, r *http.Request){
	logger.Info.Println("intialized register router........")
	var message interface{}
	message = "hello!"  
	var data = make(map[string]interface{})
	data["data"] = message
	utility.Respond(w, data)
		
}