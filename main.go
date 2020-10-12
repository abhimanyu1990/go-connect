package main

import(
	"net/http"
	"fmt"
	"go-connect/app/conf"
	"reflect"
)

func main(){
	
	// Initializing all the routes
	routes := conf.Router()
	
	err := http.ListenAndServe(":"+"8000", conf.Router()) 
	if err != nil {
		fmt.Print(err)
	}
}