package controllers

import (
	"net/http"
	"fmt"
)

var Register = func(w http.ResponseWriter, r *http.Request){
	logger.Info.Println("intialized register router")
    fmt.Fprintf(w, "Hello!")
}