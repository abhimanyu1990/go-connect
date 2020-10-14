package controllers

import (
	"net/http"
	"go-connect/app/utility"
	"github.com/gorilla/context"
)

var Register = func(w http.ResponseWriter, r *http.Request){ 
	logger.Info.Println("test...")
	body := context.Get(r,"reqBody")
	var data = make(map[string]interface{})
	data["Data"] = body
	utility.Respond(w, data)
		
}