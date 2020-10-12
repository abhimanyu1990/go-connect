package controllers

import (
	"net/http"
	"fmt"
)


var Register = func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
}